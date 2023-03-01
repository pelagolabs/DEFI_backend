package events

import (
	"context"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"io"
	"time"
	"veric-backend/internal/log"
	"veric-backend/logic/blockchain/eth"
	"veric-backend/logic/db"
)

const maxFetchLogsOnce = 50

type EventProcess interface {
	io.Closer

	Start(rt ratelimit.Limiter)
}

type Events struct {
	conf *db.Blockchain

	ethClient *eth.Client
	allCloser []io.Closer
}

func NewEvents(conf *db.Blockchain) (*Events, error) {
	e := &Events{
		conf: conf,
	}

	err := e.initClient()
	if err != nil {
		return nil, err
	}

	err = e.initProcess()
	if err != nil {
		return nil, err
	}

	err = e.initListen()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Events) initClient() error {
	var err error

	e.ethClient, err = eth.NewClient(e.conf.Endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (e *Events) initListen() error {
	rt := ratelimit.New(15)

	currencies, err := db.FindAllCurrencyByChain(e.conf.ID)
	if err != nil {
		return err
	}

	for _, erc20Config := range currencies {
		var event EventProcess
		var err error

		if erc20Config.IsNative {
			event, err = NewNativeEvent(e, e.conf.ContractName, erc20Config.Symbol)
		} else {
			event, err = NewErc20Event(e, e.conf.ContractName, erc20Config.Symbol, erc20Config.ContractAddress)
		}
		if err != nil {
			return err
		}

		event.Start(rt)
		e.allCloser = append(e.allCloser, event)
	}
	return nil
}

func (e *Events) filterRange(ctx context.Context, rt ratelimit.Limiter, watchStartFlag int64, f func(watchStart uint64, watchEnd uint64) error) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	var (
		initWatchStart       bool
		watchStart, watchEnd uint64
		err                  error
	)

	fastNext := true
	for {
		if fastNext {
			fastNext = false
		} else {
			select {
			case <-ticker.C:

			case <-ctx.Done():
				return
			}
		}

		if !initWatchStart {
			if watchStartFlag >= 0 {
				watchStart = uint64(watchStartFlag)
			} else {
				watchStart, err = e.ethClient.BlockNumber(context.Background())
				if err != nil {
					log.GetLogger().Warn("filterRange get BlockNumber err", zap.Error(err))
					continue
				}
			}
			initWatchStart = true
		}

		rt.Take()
		watchEnd, err = e.ethClient.BlockNumber(context.Background())
		if err != nil {
			log.GetLogger().Warn("filterRange get BlockNumber err", zap.Error(err))
			continue
		}

		if watchStart >= watchEnd {
			continue
		}

		if watchEnd-watchStart > maxFetchLogsOnce {
			watchEnd = watchStart + maxFetchLogsOnce
			fastNext = true
		}

		err = f(watchStart, watchEnd)
		if err != nil {
			log.GetLogger().Warn("filterRange err", zap.Error(err))
			continue
		}

		watchStart = watchEnd + 1
	}
}

func (e *Events) initProcess() error {
	return nil
}

func (e *Events) Close() {

}
