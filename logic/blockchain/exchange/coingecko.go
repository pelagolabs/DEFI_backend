package exchange

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"math/big"
	"net/http"
	"strings"
	"veric-backend/internal/log"
)

type coingeckoResponseItem struct {
	Usd float64 `json:"usd"`
}

type coingeckoResponse map[string]*coingeckoResponseItem

type coingecko struct {
}

func (c *coingecko) QueryExchanges(names []string) map[string]*big.Float {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", strings.Join(names, ","))
	resp, err := http.Get(url)
	if err != nil {
		log.GetLogger().Warn("[coingecko] query fail", zap.Error(err))
		return nil
	}
	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.GetLogger().Warn("[coingecko] read fail", zap.Error(err))
		return nil
	}

	coingeckoResp := make(coingeckoResponse)
	err = json.Unmarshal(jsonBody, &coingeckoResp)
	if err != nil {
		log.GetLogger().Warn("[coingecko] json decode fail", zap.Error(err), zap.ByteString("body", jsonBody))
		return nil
	}

	ret := make(map[string]*big.Float)
	for k, item := range coingeckoResp {
		ret[k] = big.NewFloat(item.Usd).SetPrec(64)
	}
	return ret
}
