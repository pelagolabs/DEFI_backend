package exchange

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"math/big"
	"net/http"
	"sync"
	"veric-backend/internal/log"
)

type coincapResponse struct {
	Data struct {
		Id                string `json:"id"`
		Rank              string `json:"rank"`
		Symbol            string `json:"symbol"`
		Name              string `json:"name"`
		Supply            string `json:"supply"`
		MarketCapUsd      string `json:"marketCapUsd"`
		VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
		PriceUsd          string `json:"priceUsd"`
		ChangePercent24Hr string `json:"changePercent24Hr"`
		Vwap24Hr          string `json:"vwap24Hr"`
		Explorer          string `json:"explorer"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type coincap struct {
}

func (c *coincap) QueryExchanges(names []string) map[string]*big.Float {
	ret := make(map[string]*big.Float)
	lock := &sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()

			url := fmt.Sprintf("https://api.coincap.io/v2/assets/%s", name)
			resp, err := http.Get(url)
			if err != nil {
				log.GetLogger().Warn("[coincap] query fail", zap.Error(err))
				return
			}
			defer resp.Body.Close()

			jsonBody, err := io.ReadAll(resp.Body)
			if err != nil {
				log.GetLogger().Warn("[coincap] read fail", zap.Error(err))
				return
			}

			coingeckoResp := &coincapResponse{}
			err = json.Unmarshal(jsonBody, coingeckoResp)
			if err != nil {
				log.GetLogger().Warn("[coincap] json decode fail", zap.Error(err), zap.ByteString("body", jsonBody))
				return
			}

			val, _, err := new(big.Float).SetPrec(64).Parse(coingeckoResp.Data.PriceUsd, 10)
			if err != nil {
				log.GetLogger().Warn("[coincap] parse price fail", zap.Error(err), zap.ByteString("body", jsonBody))
				return
			}

			lock.Lock()
			defer lock.Unlock()

			ret[name] = val
		}(name)
	}

	wg.Wait()
	return ret
}
