package administration

import (
	"time"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/exchange"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/http/http_util"
)

func GetGlobalOverview(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	tz := r.QueryWithDefaultInt("tz", 0)

	// get admin
	_, jwtErr := r.GetAdminJwt("admin-auth-token")
	if jwtErr != nil {
		return nil, jwtErr
	}

	nowTime := time.Now().Add(time.Duration(tz) * time.Hour)
	dateToday := nowTime.Format(util.DateLayer)
	overview := &globalOverviewResponse{
		TotalMerchantCount:     0,
		TodayNewMerchantCount:  0,
		TotalRevenueAmount:     0,
		TodayRevenueAmount:     0,
		AvailableBalanceAmount: 0,
	}

	// get merchant related
	overview.TotalMerchantCount = db.GetMerchantCountByStatus("all")
	overview.TodayNewMerchantCount = db.GetMerchantCountByStatusCreatedAt("all", util.GetTimeStrFixedByTZ(dateToday+" 00:00:00", tz), util.GetTimeStrFixedByTZ(dateToday+" 23:59:59", tz))

	// get revenue related
	feeStats, feeErr := db.FindPaymentFeeWithCurrencyByType([]string{"Platform"})
	if feeErr != nil {
		return nil, feeErr
	}

	for _, feeItem := range feeStats {
		paymentAmountRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(feeItem.TotalAmount.RawBigInt(), feeItem.Currency)
		if exchangeErr != nil {
			continue
		}

		feeAmount, _ := paymentAmountRaw.Float64()
		overview.TotalRevenueAmount += feeAmount

		if feeItem.Date.Format(util.DateLayer) == dateToday {
			overview.TodayRevenueAmount += feeAmount
		}
	}

	// get available balance
	fees, balanceErr := db.FindAllAvailablePaymentFeesByFeeType([]string{"Platform"}, db.WithPreload("Currency"))
	if balanceErr != nil {
		return nil, balanceErr
	}

	balanceAmountSet := make(map[uint]*balanceAmountItem, 0)
	for _, fee := range fees {
		if fee.FeeType == "Pool" {
			continue
		}

		canClaimedAmount := fee.TotalAmount.Copy().Sub(fee.ClaimedAmount).Sub(fee.FrozenAmount)
		if canClaimedAmount.IsZero() {
			continue
		}

		if _, ok := balanceAmountSet[fee.CurrencyId]; !ok {
			balanceAmountSet[fee.CurrencyId] = &balanceAmountItem{
				Currency:                fee.Currency,
				TotalFeeRemainingAmount: types.NewBigIntZero(),
			}
		}

		balanceAmountSet[fee.CurrencyId].TotalFeeRemainingAmount.Add(canClaimedAmount)
	}

	for _, balanceItem := range balanceAmountSet {
		balanceRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(balanceItem.TotalFeeRemainingAmount.RawBigInt(), balanceItem.Currency)
		if exchangeErr != nil {
			continue
		}

		balance, _ := balanceRaw.Float64()
		overview.AvailableBalanceAmount += balance
	}

	return overview, nil
}
