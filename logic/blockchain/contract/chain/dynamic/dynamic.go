package dynamic

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"veric-backend/internal/log"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/contract/common"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/blockchain/eth/contracts/lp_manager"
	"veric-backend/logic/blockchain/eth/contracts/payment_vault"
	"veric-backend/logic/config"
)

type DynamicChain struct {
	endpoint string
	name     string

	client         *eth.Client
	manage         *eth.TransitionManage
	saver          *common.ContractSaver
	contractPriKey *eth.PrivateKey
	conf           *config.ContractConf

	addRewardsProcess eth.TransitionProcessCaller[*AddRewardsInfo]
}

func NewDynamicChain(conf *config.ContractConf, name, endpoint string) (c *DynamicChain, err error) {
	c = &DynamicChain{
		endpoint: endpoint,
		name:     name,
		conf:     conf,
		saver:    &common.ContractSaver{Chain: "Dynamic" + name},
	}

	c.contractPriKey, err = eth.NewPrivateKey(config.Get().AccountPriKey.Contract)
	if err != nil {
		return nil, err
	}

	c.client, err = eth.NewClient(endpoint)
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

func (c *DynamicChain) initProcess() (err error) {
	{
		c.addRewardsProcess, _ = eth.NewTransitionProcessManage[*AddRewardsInfo](c.manage, c.saver, "AddRewards").
			FirstStep("Withdraw", func(i *AddRewardsInfo, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				vault, err := payment_vault.NewPaymentVault(i.WalletAddress, c.client.Client())
				if err != nil {
					return nil, err
				}

				if i.AddRewardsNative {
					return vault.WithdrawEth(opts, c.contractPriKey.Address(), i.AddRewardsAmount)
				} else {
					return vault.Withdraw(opts, i.CurrencyAddress, c.contractPriKey.Address(), i.AddRewardsAmount)
				}
			}).
			NextStep("AddRewards", func(i *AddRewardsInfo, tx *types.Transaction, receipt *types.Receipt, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				if receipt.Status == 0 {
					log.GetLogger().Error("withdraw failed", zap.String("hash", receipt.TxHash.String()))

					incidentKey := fmt.Sprintf("pool reward withdraw failed! - %s", i.BlockchainName)
					message := fmt.Sprintf(util.PoolRewardWithdrawFailedMessage, i.BlockchainName, i.AddRewardsNative, i.CurrencyAddress, i.AddRewardsAmount.Int64(), receipt.TxHash.String())
					severity := "error"

					pagerErr := util.SendPagerNotify(incidentKey, "trigger", message, severity)
					if pagerErr != nil {
						log.GetLogger().Error("[contract Withdraw] send pagerduty incident failed", zap.Error(pagerErr))
					}

					return nil, err
				}

				manager, err := lp_manager.NewLpManager(i.LpAddress, c.client.Client())
				if err != nil {
					return nil, err
				}

				if i.AddRewardsNative {
					opts.Value = i.AddRewardsAmount
					return manager.AddEthRewards(opts)
				} else {
					return manager.AddRewards(opts, i.CurrencyAddress, i.AddRewardsAmount)
				}
			}).
			LastStep("Confirm", func(i *AddRewardsInfo, tx *types.Transaction, receipt *types.Receipt) (err error) {
				if receipt.Status == 0 {
					log.GetLogger().Error("addRewards failed", zap.String("hash", receipt.TxHash.String()))

					incidentKey := fmt.Sprintf("pool reward add failed! - %s", i.BlockchainName)
					message := fmt.Sprintf(util.PoolRewardAddFailedMessage, i.BlockchainName, i.AddRewardsNative, i.CurrencyAddress, i.AddRewardsAmount.Int64(), receipt.TxHash.String())
					severity := "error"

					pagerErr := util.SendPagerNotify(incidentKey, "trigger", message, severity)
					if pagerErr != nil {
						log.GetLogger().Error("[contract AddRewards] send pagerduty incident failed", zap.Error(pagerErr))
					}
				}

				return nil
			})
	}

	return nil
}

func (c *DynamicChain) AddRewardsProcess() eth.TransitionProcessCaller[*AddRewardsInfo] {
	return c.addRewardsProcess
}
