package exchange

import (
	"fmt"
	"go.uber.org/zap"
	"math/big"
	"sync/atomic"
	"time"
	"veric-backend/internal/log"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
)

var DefaultManage *Manage

func init() {
	log.GetLogger().Info("init exchange manage...")

	manage, err := NewManage(time.Minute)
	if err != nil {
		panic(err)
	}

	DefaultManage = manage
}

type exchangeApi interface {
	QueryExchanges(names []string) map[string]*big.Float
}

type Manage struct {
	updateInterval time.Duration
	apis           []exchangeApi
	exchangeMap    atomic.Pointer[map[string]*big.Float]
}

func NewManage(updateInterval time.Duration) (*Manage, error) {
	m := &Manage{
		updateInterval: updateInterval,
		apis: []exchangeApi{
			&coingecko{},
			&coincap{},
		},
	}
	err := m.updateOnce()
	if err != nil {
		return nil, err
	}

	go m.update()

	return m, nil
}

func (m *Manage) update() {
	for range time.Tick(m.updateInterval) {
		err := m.updateOnce()
		if err != nil {
			log.GetLogger().Warn("[exchange manage] update fail", zap.Error(err))
		}
	}
}

func (m *Manage) updateOnce() error {
	currencies, err := db.FindAvailableCurrencies(db.WithPreload("Blockchain"))
	if err != nil {
		return err
	}

	names := make([]string, 0, len(currencies))
	priceMap := make(map[string]*big.Float)
	count := 0
	for _, currency := range currencies {
		names = append(names, currency.SearchName)
		priceMap[currency.SearchName] = nil
		count++
	}

	for _, api := range m.apis {
		if count == 0 {
			break
		}

		exchanges := api.QueryExchanges(names)
		if exchanges == nil || len(exchanges) == 0 {
			continue
		}
		names = names[:0]
		count = 0
		for _, currency := range currencies {
			if priceMap[currency.SearchName] == nil {
				if exchange, ok := exchanges[currency.SearchName]; ok && exchange != nil {
					priceMap[currency.SearchName] = exchange
				} else {
					names = append(names, currency.SearchName)
					count++
				}
			}
		}
	}

	m.exchangeMap.Store(&priceMap)
	return nil
}

func (m *Manage) valueOfDecimalCount(count uint) *types.BigFloat {
	return types.NewBigFloatUseBigInt(
		types.NewBigIntFast(10).Pow(types.NewBigIntFast(int64(count))),
	)
}

func (m *Manage) ExchangeCoinToNaturalAmount(amount *big.Int, currency *db.Currency) *big.Float {
	// amount / (10^currency.DecimalCount)
	return types.NewBigFloatUseBigInt(types.NewBigInt(amount)).
		Div(m.valueOfDecimalCount(currency.DecimalCount)).
		RawBigFloat()
}

func (m *Manage) ExchangeCoinToUSD(amount *big.Int, currency *db.Currency) (*big.Float, error) {
	exchangeMap := *(m.exchangeMap.Load())
	if exchange, ok := exchangeMap[currency.SearchName]; ok && exchange != nil {
		// amount / (10^currency.DecimalCount) * exchangeMap[currency.SearchName]
		decimal := types.NewBigIntFast(10).Pow(types.NewBigIntFast(int64(currency.DecimalCount)))
		result := types.
			NewBigFloatUseBigInt(types.NewBigInt(amount)).
			Div(types.NewBigFloatUseBigInt(decimal)).
			Mul(types.NewBigFloat(exchange))

		return result.RawBigFloat(), nil
	} else {
		return nil, fmt.Errorf("not found %s exchange", currency.Name)
	}
}

func (m *Manage) ExchangeUSDToCoin(amount *big.Float, currency *db.Currency) (*big.Int, error) {
	exchangeMap := *(m.exchangeMap.Load())
	if exchange, ok := exchangeMap[currency.SearchName]; ok && exchange != nil {
		// amount / exchangeMap[currency.SearchName] * (10^currency.DecimalCount)
		decimal := types.NewBigIntFast(10).Pow(types.NewBigIntFast(int64(currency.DecimalCount)))
		result := types.
			NewBigFloat(amount).
			Div(types.NewBigFloat(exchange)).
			Mul(types.NewBigFloatUseBigInt(decimal)).
			RoundToInt().
			RawBigInt()

		return result, nil
	} else {
		return nil, fmt.Errorf("not found %s exchange", currency.Name)
	}
}

func (m *Manage) QueryPrice(currency *db.Currency) (*big.Float, error) {
	exchangeMap := *(m.exchangeMap.Load())
	if exchange, ok := exchangeMap[currency.SearchName]; ok && exchange != nil {
		return exchange, nil
	} else {
		return nil, fmt.Errorf("not found %s exchange", currency.Name)
	}
}
