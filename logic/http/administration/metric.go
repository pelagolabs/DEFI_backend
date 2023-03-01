package administration

import (
	"net/http"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/http/http_util"
)

func GetCurrencyMaxBalanceList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	secret := r.QueryWithDefault("secret", "")
	if secret == "" || secret != config.Get().HTTP.MetricSecret {
		return nil, http_util.NewHttpError(http.StatusNotFound, "404 page not found")
	}

	// get all available currency
	currencies, getCurrencyErr := db.FindAvailableNotHiddenCurrencies(db.WithPreload("Blockchain"))
	if getCurrencyErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get all available currency failed")
	}

	// get all merchant claimable balance
	balances, getBalanceErr := db.FindAllUnClaimedPaymentBalance()
	if getBalanceErr != nil {
		return nil, http_util.NewHttpError(http.StatusInternalServerError, "get all merchant balance failed")
	}

	merchantBalanceSet := make(map[uint]map[uint]*types.BigInt, 0)
	currencyMaxBalanceSet := make(map[uint]*types.BigInt, 0)
	for _, balance := range balances {
		if _, ok := merchantBalanceSet[balance.ApplicationId]; !ok {
			merchantBalanceSet[balance.ApplicationId] = make(map[uint]*types.BigInt, 0)
		}

		if _, ok := merchantBalanceSet[balance.ApplicationId][balance.CurrencyId]; !ok {
			merchantBalanceSet[balance.ApplicationId][balance.CurrencyId] = types.NewBigIntZero()
		}

		claimableBalance := balance.TotalAmount.Copy().Sub(balance.ClaimedAmount)
		merchantBalanceSet[balance.ApplicationId][balance.CurrencyId].Add(claimableBalance)
	}

	// compare to find max balance
	for _, balanceSet := range merchantBalanceSet {
		for currencyId, balance := range balanceSet {
			if _, ok := currencyMaxBalanceSet[currencyId]; !ok {
				currencyMaxBalanceSet[currencyId] = types.NewBigIntZero()
			}

			if balance.Cmp(currencyMaxBalanceSet[currencyId]) == 1 {
				currencyMaxBalanceSet[currencyId] = balance.Copy()
			}
		}
	}

	balanceList := make([]*currencyMaxBalanceListItem, 0)
	for _, currency := range currencies {
		if val, ok := currencyMaxBalanceSet[currency.ID]; !ok {
			balanceList = append(balanceList, &currencyMaxBalanceListItem{
				Currency:   &currency,
				MaxBalance: "0",
			})
		} else {
			balanceList = append(balanceList, &currencyMaxBalanceListItem{
				Currency:   &currency,
				MaxBalance: val.String(),
			})
		}
	}

	return &currencyMaxBalanceListResponse{
		Total:      uint(len(currencies)),
		Currencies: balanceList,
	}, nil
}
