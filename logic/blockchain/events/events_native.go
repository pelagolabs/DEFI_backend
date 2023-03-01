package events

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"math/big"
	"sync/atomic"
	"time"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/blockchain/eth/contracts/payment_vault"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/tasks"
)

type NativeEvent struct {
	ev *Events

	ChainName  string
	CoinName   string
	BlockStart int64

	isStart atomic.Bool
	isClose atomic.Bool
	ctx     context.Context
	cancel  context.CancelFunc

	sol      *payment_vault.PaymentVault
	currency *db.Currency
	solAbi   *abi.ABI
}

func NewNativeEvent(ev *Events, chainName, coinName string) (e *NativeEvent, err error) {
	e = &NativeEvent{
		ev:         ev,
		ChainName:  chainName,
		CoinName:   coinName,
		BlockStart: -1,
	}

	e.ctx, e.cancel = context.WithCancel(context.Background())
	e.sol, err = payment_vault.NewPaymentVault(common.Address{}, e.ev.ethClient.Client())
	if err != nil {
		return nil, err
	}
	e.solAbi, err = payment_vault.PaymentVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *NativeEvent) Start(rt ratelimit.Limiter) {
	if e.isStart.CompareAndSwap(false, true) {
		transferChan := make(chan *payment_vault.PaymentVaultReceivedEth, 100)
		go func() {
			defer close(transferChan)

			e.refreshStart()

			e.ev.filterRange(e.ctx, rt, e.BlockStart, func(watchStart, watchEnd uint64) error {
				topics, err := abi.MakeTopics([]interface{}{e.solAbi.Events["ReceivedEth"].ID})
				if err != nil {
					return err
				}

				logs, err := e.ev.ethClient.FilterLogs(context.Background(), ethereum.FilterQuery{
					Topics:    topics,
					FromBlock: new(big.Int).SetUint64(watchStart),
					ToBlock:   new(big.Int).SetUint64(watchEnd),
				})
				if err != nil {
					return err
				}

				if e.currency != nil {
					e.currency.EventBlock = watchStart
					err = db.UpdateCurrencyEventBlock(e.currency.ID, e.currency.EventBlock)
					if err != nil {
						log.GetLogger().Warn("[NativeEvent] update db fail", zap.Error(err))
					}
				}

				for _, rawLog := range logs {
					event, err := e.sol.ParseReceivedEth(rawLog)
					if err != nil {
						return err
					}

					transferChan <- event
				}

				return nil
			})
		}()

		go func() {
			e.processTransfers(transferChan)
			e.isClose.Store(true)
		}()
	}
}

func (e *NativeEvent) processTransfers(transferChan chan *payment_vault.PaymentVaultReceivedEth) {
	for transfer := range transferChan {
		if !contract.DefaultContract.IsContractWalletAddress(e.ChainName, transfer.Raw.Address) {
			continue
		}

		err := e.processTransfer(transfer)
		if err != nil {
			log.GetLogger().Warn("[processTransfers]", zap.Any("event", transfer), zap.Error(err))
		}
	}
}

func (e *NativeEvent) processTransfer(ev *payment_vault.PaymentVaultReceivedEth) error {
	tx, err := db.QueryPaymentTxByHash(fmt.Sprintf("%s:%d", ev.Raw.TxHash.String(), ev.Raw.Index))
	if err != nil {
		return err
	}

	if tx.ID > 0 {
		return nil
	}

	config, err := e.toPaymentIncomeTaskConfig(ev)
	if err != nil {
		return err
	}

	return tasks.DefaultManage.NewTask(tasks.TaskTypePaymentIncome, config)
}

func (e *NativeEvent) toPaymentIncomeTaskConfig(ev *payment_vault.PaymentVaultReceivedEth) (*tasks.PaymentIncomeTaskConfig, error) {
	raw, err := json.Marshal(ev)
	if err != nil {
		return nil, err
	}

	return &tasks.PaymentIncomeTaskConfig{
		From:      ev.Who,
		To:        ev.Raw.Address,
		Amount:    types.NewBigInt(ev.Amount),
		UniqueKey: fmt.Sprintf("%s:%d", ev.Raw.TxHash.String(), ev.Raw.Index),
		Raw:       raw,
	}, nil
}

func (e *NativeEvent) Close() error {
	e.cancel()
	for !e.isClose.Load() {
		time.Sleep(16 * time.Millisecond)
	}

	return nil
}

func (e *NativeEvent) refreshStart() {
	blockchain, err := db.FindBlockchainByContractName(e.ChainName)
	if blockchain.ID == 0 || err != nil {
		log.GetLogger().Warn("ERC20 State: can not find chain", zap.String("name", e.ChainName))
		return
	}

	currency, err := db.FindCurrencyByChainAndSymbol(blockchain.ID, e.CoinName)
	if currency.ID == 0 || err != nil {
		log.GetLogger().Warn("ERC20 State: can not find currency", zap.String("name", e.CoinName))
		return
	}

	e.currency = currency
	if currency.EventBlock > 0 {
		e.BlockStart = int64(currency.EventBlock)
	}
}
