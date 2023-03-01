package events

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"sync/atomic"
	"time"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/blockchain/eth/contracts/erc20"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/tasks"
)

type Erc20Event struct {
	ev *Events

	ChainName  string
	CoinName   string
	Addr       common.Address
	BlockStart int64

	isStart atomic.Bool
	isClose atomic.Bool
	ctx     context.Context
	cancel  context.CancelFunc

	sol      *erc20.Erc20
	currency *db.Currency
}

func NewErc20Event(ev *Events, chainName, coinName string, addr string) (e *Erc20Event, err error) {
	e = &Erc20Event{
		ev:         ev,
		ChainName:  chainName,
		CoinName:   coinName,
		Addr:       common.HexToAddress(addr),
		BlockStart: -1,
	}

	e.ctx, e.cancel = context.WithCancel(context.Background())
	e.sol, err = erc20.NewErc20(e.Addr, e.ev.ethClient.Client())
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Erc20Event) Start(rt ratelimit.Limiter) {
	if e.isStart.CompareAndSwap(false, true) {
		transferChan := make(chan *erc20.Erc20Transfer, 100)
		go func() {
			defer close(transferChan)

			e.refreshStart()

			e.ev.filterRange(e.ctx, rt, e.BlockStart, func(watchStart, watchEnd uint64) error {
				transferEvent, err := e.sol.FilterTransfer(&bind.FilterOpts{Start: watchStart, End: &watchEnd}, nil, nil)
				if err != nil {
					return err
				}
				defer transferEvent.Close()

				if e.currency != nil {
					e.currency.EventBlock = watchStart
					err = db.UpdateCurrencyEventBlock(e.currency.ID, e.currency.EventBlock)
					if err != nil {
						log.GetLogger().Warn("[Erc20Event] update db fail", zap.Error(err))
					}
				}

				for transferEvent.Next() {
					if transferEvent.Error() != nil {
						return transferEvent.Error()
					}
					transferChan <- transferEvent.Event
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

func (e *Erc20Event) processTransfers(transferChan chan *erc20.Erc20Transfer) {
	for transfer := range transferChan {
		if !contract.DefaultContract.IsContractWalletAddress(e.ChainName, transfer.To) {
			continue
		}

		err := e.processTransfer(transfer)
		if err != nil {
			log.GetLogger().Warn("[processTransfers]", zap.Any("event", transfer), zap.Error(err))
		}
	}
}

func (e *Erc20Event) processTransfer(ev *erc20.Erc20Transfer) error {
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

func (e *Erc20Event) toPaymentIncomeTaskConfig(ev *erc20.Erc20Transfer) (*tasks.PaymentIncomeTaskConfig, error) {
	raw, err := json.Marshal(ev)
	if err != nil {
		return nil, err
	}

	return &tasks.PaymentIncomeTaskConfig{
		From:      ev.From,
		To:        ev.To,
		Amount:    types.NewBigInt(ev.Value),
		UniqueKey: fmt.Sprintf("%s:%d", ev.Raw.TxHash.String(), ev.Raw.Index),
		Raw:       raw,
	}, nil
}

func (e *Erc20Event) Close() error {
	e.cancel()
	for !e.isClose.Load() {
		time.Sleep(16 * time.Millisecond)
	}

	return nil
}

func (e *Erc20Event) refreshStart() {
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
