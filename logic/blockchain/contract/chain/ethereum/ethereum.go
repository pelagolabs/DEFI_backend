package ethereum

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"veric-backend/logic/blockchain/contract/common"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/blockchain/eth/contracts/erc20"
	"veric-backend/logic/blockchain/eth/contracts/lp_manager"
	"veric-backend/logic/blockchain/eth/contracts/multicall"
	"veric-backend/logic/blockchain/eth/contracts/payment_vault"
	"veric-backend/logic/blockchain/eth/contracts/token_stake"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
)

type EthereumChain struct {
	conf           *config.ContractConf
	client         *eth.Client
	manage         *eth.TransitionManage
	saver          *common.ContractSaver
	contractPriKey *eth.PrivateKey

	withdrawProcess       eth.TransitionProcessCaller[*WithdrawInfo]
	withdrawNativeProcess eth.TransitionProcessCaller[*WithdrawInfo]
	lpSwitchProcess       eth.TransitionProcessCaller[*LpSwitchInfo]
	farmSwitchProcess     eth.TransitionProcessCaller[*FarmSwitchInfo]
}

func NewEthereumChain(conf *config.ContractConf) (c *EthereumChain, err error) {
	c = &EthereumChain{
		conf:  conf,
		saver: &common.ContractSaver{Chain: "Ethereum"},
	}

	c.contractPriKey, err = eth.NewPrivateKey(config.Get().AccountPriKey.Contract)
	if err != nil {
		return nil, err
	}

	c.client, err = eth.NewClient(c.conf.EthereumEndpoint)
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

func (c *EthereumChain) initProcess() (err error) {

	{
		c.withdrawProcess, _ = eth.NewTransitionProcessManage[*WithdrawInfo](c.manage, c.saver, "Withdraw").
			FirstStep("Transition", func(i *WithdrawInfo, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				vault, err := payment_vault.NewPaymentVault(i.WalletAddress, c.client.Client())
				if err != nil {
					return nil, err
				}

				next, err = vault.Withdraw(opts, i.CurrencyAddress, i.WithdrawAddress, i.WithdrawAmount)
				if i.WithdrawId > 0 {
					withdraw, dbErr := db.FindWithdrawById(i.WithdrawId)
					if dbErr != nil {
						return next, err
					}

					withdraw.TxHash = next.Hash().String()
					dbErr = db.SaveWithdraw(withdraw)
					if dbErr != nil {
						return next, err
					}
				}

				return next, err
			}).
			LastStep("Confirm", func(i *WithdrawInfo, tx *types.Transaction, receipt *types.Receipt) (err error) {
				if i.WithdrawId > 0 {
					withdraw, err := db.FindWithdrawById(i.WithdrawId)
					if err != nil {
						return err
					}
					if receipt.Status == 1 {
						withdraw.Status = db.WithdrawStatusComplete
					} else {
						withdraw.Status = db.WithdrawStatusClosed
					}
					withdraw.TxHash = receipt.TxHash.String()
					return db.SaveWithdraw(withdraw)
				} else {
					return nil
				}
			})
	}

	{
		c.withdrawNativeProcess, _ = eth.NewTransitionProcessManage[*WithdrawInfo](c.manage, c.saver, "Withdraw").
			FirstStep("Transition", func(i *WithdrawInfo, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				vault, err := payment_vault.NewPaymentVault(i.WalletAddress, c.client.Client())
				if err != nil {
					return nil, err
				}

				next, err = vault.WithdrawEth(opts, i.WithdrawAddress, i.WithdrawAmount)
				if i.WithdrawId > 0 {
					withdraw, dbErr := db.FindWithdrawById(i.WithdrawId)
					if dbErr != nil {
						return next, err
					}

					withdraw.TxHash = next.Hash().String()
					dbErr = db.SaveWithdraw(withdraw)
					if dbErr != nil {
						return next, err
					}
				}

				return next, err
			}).
			LastStep("Confirm", func(i *WithdrawInfo, tx *types.Transaction, receipt *types.Receipt) (err error) {
				if i.WithdrawId > 0 {
					withdraw, err := db.FindWithdrawById(i.WithdrawId)
					if err != nil {
						return err
					}
					if receipt.Status == 1 {
						withdraw.Status = db.WithdrawStatusComplete
					} else {
						withdraw.Status = db.WithdrawStatusClosed
					}
					withdraw.TxHash = receipt.TxHash.String()
					return db.SaveWithdraw(withdraw)
				} else {
					return nil
				}
			})
	}

	{
		c.lpSwitchProcess, _ = eth.NewTransitionProcessManage[*LpSwitchInfo](c.manage, c.saver, "LpSwitch").
			FirstStep("Transition", func(i *LpSwitchInfo, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				manager, err := lp_manager.NewLpManager(i.LpAddress, c.client.Client())
				if err != nil {
					return nil, err
				}

				if i.IsGlobal {
					if i.Action == "pause" {
						return manager.Pause(opts)
					} else {
						return manager.Unpause(opts)
					}
				} else {
					if i.IsNativeCurrency {
						if i.Action == "pause" {
							return manager.PauseEth(opts)
						} else {
							return manager.UnpauseEth(opts)
						}
					} else {
						if i.Action == "pause" {
							return manager.PauseToken(opts, i.CurrencyAddress)
						} else {
							return manager.UnpauseToken(opts, i.CurrencyAddress)
						}
					}
				}

			}).
			LastStep("Confirm", func(i *LpSwitchInfo, tx *types.Transaction, receipt *types.Receipt) (err error) {
				return nil
			})
	}

	{
		c.farmSwitchProcess, _ = eth.NewTransitionProcessManage[*FarmSwitchInfo](c.manage, c.saver, "FarmSwitch").
			FirstStep("Transition", func(i *FarmSwitchInfo, opts *bind.TransactOpts) (next *types.Transaction, err error) {
				farmManager, err := token_stake.NewTokenStake(i.FarmAddress, c.client.Client())
				if err != nil {
					return nil, err
				}

				if i.Action == "pause" {
					return farmManager.Pause(opts)
				} else {
					return farmManager.Unpause(opts)
				}
			}).
			LastStep("Confirm", func(i *FarmSwitchInfo, tx *types.Transaction, receipt *types.Receipt) (err error) {
				return nil
			})
	}

	return nil
}

func (c *EthereumChain) QueryEnoughWalletsMaxAmount(wallets []*db.WalletPool, allowLP bool, currency *db.Currency, needAmount *big.Int) (wallet string, err error) {
	factory := multicall.NewMultiCallFactory(ethCommon.HexToAddress(c.conf.MultiCallAddr), c.client.Client())
	var maxWallet *db.WalletPool

	maxWalletAmount := big.NewInt(0)
	for _, wallet := range wallets {
		w := wallet
		err = factory.PushCall(multicall.MultiCallMethod{
			ContractAddress:  ethCommon.HexToAddress(currency.ContractAddress),
			ContractMetadata: erc20.Erc20MetaData,
			MethodName:       "balanceOf",
			Params:           []any{ethCommon.HexToAddress(w.Addr)},
			ResultProcess: func(result []any) error {
				val := abi.ConvertType(result[0], new(big.Int)).(*big.Int)

				if val.Cmp(maxWalletAmount) > 0 {
					maxWalletAmount = val
					maxWallet = w
				}
				return nil
			},
		})
		if err != nil {
			return "", err
		}
	}

	lpAmount := big.NewInt(0)
	lpAddr := ""
	if allowLP {
		pool, err := db.FindPoolByCurrencyId(currency.ID)
		if err == nil {
			manager, err := lp_manager.NewLpManager(ethCommon.HexToAddress(pool.Address), c.client.Client())
			if err != nil {
				return "", err
			}

			tokenAddress, err := manager.GetPool(nil, ethCommon.HexToAddress(currency.ContractAddress))
			if err != nil {
				return "", err
			}

			_ = factory.PushCall(multicall.MultiCallMethod{
				ContractAddress:  ethCommon.HexToAddress(currency.ContractAddress),
				ContractMetadata: erc20.Erc20MetaData,
				MethodName:       "balanceOf",
				Params:           []any{tokenAddress},
				ResultProcess: func(result []any) error {
					lpAddr = pool.Address
					lpAmount = abi.ConvertType(result[0], new(big.Int)).(*big.Int)
					return nil
				},
			})
		}
	}

	err = factory.DoRead(nil)
	if err != nil {
		return "", err
	}

	if maxWalletAmount.Cmp(needAmount) >= 0 {
		return maxWallet.Addr, nil
	} else if lpAmount.Cmp(needAmount) >= 0 {
		return lpAddr, nil
	} else {
		return "", errors.New("all wallet balance is not enough")
	}
}

func (c *EthereumChain) QueryEnoughWalletsMaxNativeAmount(wallets []*db.WalletPool, allowLP bool, currency *db.Currency, needAmount *big.Int) (wallet string, err error) {
	factory := multicall.NewMultiCallFactory(ethCommon.HexToAddress(c.conf.MultiCallAddr), c.client.Client())
	var maxWallet *db.WalletPool

	maxWalletAmount := big.NewInt(0)
	for _, wallet := range wallets {
		w := wallet
		err = factory.PushCall(multicall.MultiCallMethod{
			ContractAddress:  ethCommon.HexToAddress(c.conf.MultiCallAddr),
			ContractMetadata: multicall.MulticallMetaData,
			MethodName:       "getEthBalance",
			Params:           []any{ethCommon.HexToAddress(w.Addr)},
			ResultProcess: func(result []any) error {
				val := abi.ConvertType(result[0], new(big.Int)).(*big.Int)

				if val.Cmp(maxWalletAmount) > 0 {
					maxWalletAmount = val
					maxWallet = w
				}
				return nil
			},
		})
		if err != nil {
			return "", err
		}
	}

	lpAmount := big.NewInt(0)
	lpAddr := ""
	if allowLP {
		pool, err := db.FindPoolByCurrencyId(currency.ID)
		if err == nil {
			manager, err := lp_manager.NewLpManager(ethCommon.HexToAddress(pool.Address), c.client.Client())
			if err != nil {
				return "", err
			}

			tokenAddress, err := manager.EthPool(nil)
			if err != nil {
				return "", err
			}

			_ = factory.PushCall(multicall.MultiCallMethod{
				ContractAddress:  ethCommon.HexToAddress(c.conf.MultiCallAddr),
				ContractMetadata: multicall.MulticallMetaData,
				MethodName:       "getEthBalance",
				Params:           []any{tokenAddress},
				ResultProcess: func(result []any) error {
					lpAddr = pool.Address
					lpAmount = abi.ConvertType(result[0], new(big.Int)).(*big.Int)
					return nil
				},
			})
		}
	}

	err = factory.DoRead(nil)
	if err != nil {
		return "", err
	}

	if maxWalletAmount.Cmp(needAmount) >= 0 {
		return maxWallet.Addr, nil
	} else if lpAmount.Cmp(needAmount) >= 0 {
		return lpAddr, nil
	} else {
		return "", errors.New("all wallet balance is not enough")
	}
}

func (c *EthereumChain) WithdrawProcess() eth.TransitionProcessCaller[*WithdrawInfo] {
	return c.withdrawProcess
}

func (c *EthereumChain) WithdrawNativeProcess() eth.TransitionProcessCaller[*WithdrawInfo] {
	return c.withdrawNativeProcess
}

func (c *EthereumChain) LpSwitchProcess() eth.TransitionProcessCaller[*LpSwitchInfo] {
	return c.lpSwitchProcess
}

func (c *EthereumChain) FarmSwitchProcess() eth.TransitionProcessCaller[*FarmSwitchInfo] {
	return c.farmSwitchProcess
}
