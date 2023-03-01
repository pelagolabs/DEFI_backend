package harmony

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"math/big"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/contract/common"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/blockchain/eth/contracts/address_pool"
	"veric-backend/logic/blockchain/eth/contracts/did_registry"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
)

type HarmonyChain struct {
	conf           *config.ContractConf
	client         *eth.Client
	manage         *eth.TransitionManage
	saver          *common.ContractSaver
	contractPriKey *eth.PrivateKey

	addressPoolContract *address_pool.AddressPool
	didRegistryContract *did_registry.DidRegistry

	releaseVaultProcess eth.TransitionProcessCaller[*VaultInfo]
	useVaultProcess     eth.TransitionProcessCaller[*VaultInfo]
	didRegistryProcess  eth.TransitionProcessCaller[*DidInfo]
}

func NewHarmonyChain(conf *config.ContractConf) (c *HarmonyChain, err error) {
	c = &HarmonyChain{
		conf:  conf,
		saver: &common.ContractSaver{Chain: "Harmony"},
	}

	c.contractPriKey, err = eth.NewPrivateKey(config.Get().AccountPriKey.Contract)
	if err != nil {
		return nil, err
	}

	c.client, err = eth.NewClient(c.conf.HarmonyEndpoint)
	if err != nil {
		return nil, err
	}

	chainId, err := c.client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	c.manage = eth.NewTransitionManage(c.contractPriKey, chainId, c.client, c.saver)
	err = c.initProcess()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *HarmonyChain) initProcess() (err error) {
	c.addressPoolContract, err = address_pool.NewAddressPool(ethCommon.HexToAddress(c.conf.AddressPoolAddr), c.client.Client())
	if err != nil {
		return err
	}

	c.didRegistryContract, err = did_registry.NewDidRegistry(ethCommon.HexToAddress(c.conf.DidRegistryAddr), c.client.Client())
	if err != nil {
		return err
	}

	{
		c.releaseVaultProcess, _ = eth.NewTransitionProcessManage[*VaultInfo](c.manage, c.saver, "ReleaseVault").
			FirstStep("Transition", func(i *VaultInfo, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				// limit harmony gas consume
				opts.GasPrice = big.NewInt(101000000000)

				return c.addressPoolContract.ReleaseVault(opts, i.ChainName, i.Address)
			}).
			LastStep("Confirm", func(i *VaultInfo, tx *types.Transaction, receipt *types.Receipt) (err error) {
				return nil
			})
	}

	{
		c.useVaultProcess, _ = eth.NewTransitionProcessManage[*VaultInfo](c.manage, c.saver, "UseVault").
			FirstStep("Transition", func(i *VaultInfo, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				randomId, _ := big.NewInt(0).SetString(i.RandomId, 10)
				// limit harmony gas consume
				opts.GasPrice = big.NewInt(101000000000)

				return c.addressPoolContract.UseAvailableVault(opts, randomId, i.ChainName)
			}).
			LastStep("Confirm", func(i *VaultInfo, tx *types.Transaction, receipt *types.Receipt) (err error) {
				for _, l := range receipt.Logs {
					vault, err := c.addressPoolContract.ParseUsingVault(*l)
					if err != nil {
						return err
					}

					i.Address = vault.VaultAddress
				}
				return nil
			})
	}

	{
		c.didRegistryProcess, _ = eth.NewTransitionProcessManage[*DidInfo](c.manage, c.saver, "DidRegistry").
			FirstStep("DoDidRegistry", func(i *DidInfo, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				// limit harmony gas consume
				opts.GasPrice = big.NewInt(101000000000)

				return c.didRegistryContract.RegisterDid(opts, i.Did, i.PubKey)
			}).
			LastStep("ConfirmDidRegistry", func(i *DidInfo, tx *types.Transaction, receipt *types.Receipt) (err error) {
				user, userErr := db.FindUserByDid(i.Did)
				if userErr != nil {
					log.GetLogger().Warn("[contract DidRegistryProcess] find user failed", zap.Error(userErr))
					return nil
				}

				if user == nil || user.ID == 0 {
					log.GetLogger().Warn("[contract DidRegistryProcess] can not found user", zap.String("did", i.Did))
					return nil
				}

				if user.DidUpStatus != db.DidUpStatusPending {
					return nil
				}

				user.DidUpStatus = db.DidUpStatusSuccess
				user.DidUpTx = tx.Hash().String()

				saveErr := db.SaveUser(user)
				if saveErr != nil {
					log.GetLogger().Warn("[contract DidRegistryProcess] update user failed", zap.Error(saveErr), zap.String("did", i.Did))
					return nil
				}

				return nil
			})
	}

	return nil
}

func (c *HarmonyChain) ReleaseVaultProcess() eth.TransitionProcessCaller[*VaultInfo] {
	return c.releaseVaultProcess
}

func (c *HarmonyChain) UseVaultProcess() eth.TransitionProcessCaller[*VaultInfo] {
	return c.useVaultProcess
}

func (c *HarmonyChain) DidRegistryProcess() eth.TransitionProcessCaller[*DidInfo] {
	return c.didRegistryProcess
}
