package merchant

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
	"veric-backend/internal/log"
	"veric-backend/internal/util"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
	"veric-backend/logic/http/http_util"
)

var (
	exchangeRateCache = new(util.SyncedMap[string, *exchangeRateItem])
)

func GetAvailableCurrency(r *http_util.HTTPContext) (resp interface{}, respErr error) {
	currencies, currencyErr := db.FindAvailableNotHiddenCurrencies(db.WithPreload("Blockchain"))
	if currencyErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get currency currency failed. please try again later")
	}

	return &availableCurrencyResponse{
		Currencies: currencies,
	}, nil
}

func GetExchangeRate(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	from := "USD" // current support
	to := strings.ToUpper(r.QueryWithDefault("to", "USD"))
	rate := "1"

	// get user jwt claim
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	if to != "USD" {
		cacheKey := fmt.Sprintf("%s-%s", from, to)
		nowTime := time.Now()
		if item, ok := exchangeRateCache.Load(cacheKey); ok && nowTime.Before(item.ExpiredAt) {
			rate = item.Rate
		} else {
			exchangeRate, getErr := getLatestExchangeRate(from, to)
			if getErr != nil {
				log.GetLogger().Error("get exchange rate failed.", zap.Error(getErr))
				return nil, http_util.NewHttpError(http.StatusInternalServerError, "get exchange rate failed. please try again later")
			}

			rate = exchangeRate
			newCacheItem := &exchangeRateItem{
				From:      from,
				To:        to,
				Rate:      exchangeRate,
				ExpiredAt: nowTime.Add(time.Hour),
			}

			exchangeRateCache.Store(cacheKey, newCacheItem)
		}
	}

	return &exchangeRateResponse{
		From: "USD",
		To:   to,
		Rate: rate,
	}, nil
}

func getLatestExchangeRate(from, to string) (string, error) {
	key := config.Get().TianApi.Key
	url := fmt.Sprintf("https://apis.tianapi.com/fxrate/index?key=%s&fromcoin=%s&tocoin=%s&money=1", key, from, to)
	resRaw, getErr := util.HttpGet(url)
	if getErr != nil {
		return "", getErr
	}

	resObj := &tianApiExchangeRateResponse{}
	jsonErr := json.Unmarshal(resRaw, resObj)
	if jsonErr != nil {
		return "", jsonErr
	}

	if resObj.Code != 200 {
		return "", errors.New(resObj.Msg)
	}

	return resObj.Result.Money, nil
}
