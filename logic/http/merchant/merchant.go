package merchant

import (
	"fmt"
	"net/http"
	"time"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/exchange"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/http/http_util"
)

func GetBaseSummary(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	tz := r.QueryWithDefaultInt("tz", 0)

	// get user jwt claim
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, 0)
	if err != nil {
		return nil, err
	}

	nowTime := time.Now().Add(time.Duration(tz) * time.Hour)
	dateToday := nowTime.Format(util.DateLayer)
	baseSummary := &merchantBaseSummaryResponse{
		TotalBalance:      0,
		TodayTotalPayment: 0,
	}

	// get total balance
	balanceStats, balanceErr := db.FindUnClaimedPaymentBalanceWithCurrencyByApp(application.ID)
	if balanceErr != nil {
		return nil, balanceErr
	}

	fmt.Printf("application: %+v balanceStats: %+v\n", application, balanceStats)

	merchantBalanceCurrencySet := make(map[uint]*merchantBalanceCurrencyItem, 0)

	for _, balanceItem := range balanceStats {
		if _, ok := merchantBalanceCurrencySet[balanceItem.CurrencyId]; !ok {
			merchantBalanceCurrencySet[balanceItem.CurrencyId] = &merchantBalanceCurrencyItem{
				Currency:      balanceItem.Currency,
				TotalAmount:   types.NewBigIntFast(0),
				ClaimedAmount: types.NewBigIntFast(0),
			}
		}

		merchantBalanceCurrencySet[balanceItem.CurrencyId].TotalAmount = merchantBalanceCurrencySet[balanceItem.CurrencyId].TotalAmount.Add(balanceItem.TotalAmount)
		merchantBalanceCurrencySet[balanceItem.CurrencyId].ClaimedAmount = merchantBalanceCurrencySet[balanceItem.CurrencyId].ClaimedAmount.Add(balanceItem.ClaimedAmount)
	}

	for _, balanceItem := range merchantBalanceCurrencySet {
		totalBalanceBig := balanceItem.TotalAmount.Sub(balanceItem.ClaimedAmount)
		if !totalBalanceBig.IsZero() {
			amountInUSDRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(totalBalanceBig.RawBigInt(), balanceItem.Currency)
			if exchangeErr != nil {
				continue
			}

			amountInUSD, _ := amountInUSDRaw.Float64()
			baseSummary.TotalBalance += amountInUSD
		}
	}

	// get today summary
	todayStats, sumErr := db.FindApplicationFinishPaymentWithCurrencyByTime(application.ID, util.GetTimeStrFixedByTZ(dateToday+" 00:00:00", tz), util.GetTimeStrFixedByTZ(dateToday+" 23:59:59", tz))
	if sumErr != nil {
		return nil, sumErr
	}

	for _, paymentItem := range todayStats {
		if paymentItem.CollectionAmount.IsZero() {
			continue
		}

		canWithdrawAmount := paymentItem.CollectionAmount.Copy().Sub(paymentItem.PlatformFeeAmount).Sub(paymentItem.PoolFeeAmount).RawBigInt()
		paymentAmountRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(canWithdrawAmount, paymentItem.Currency)
		if exchangeErr != nil {
			continue
		}

		paymentAmount, _ := paymentAmountRaw.Float64()
		baseSummary.TodayTotalPayment += paymentAmount
	}

	return baseSummary, nil
}

func GetPaymentStatChart(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	gap := r.QueryWithDefault("gap", "24h")
	tz := r.QueryWithDefaultInt("tz", 0)

	// get user jwt claim
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, 0)
	if err != nil {
		return nil, err
	}

	nowTime := time.Now()
	dateToday := nowTime.Format(util.DateLayer)
	dateBegin := ""
	amountStatSet := make(map[string]float64, 0)

	// init x-title
	switch gap {
	case "24h":
		timeBegin := nowTime.AddDate(0, 0, -1)
		for i := 0; i < 24; i++ {
			timeCurrent := timeBegin.Add(time.Duration(i+tz) * time.Hour).Format(util.TimeHourLayer)
			amountStatSet[timeCurrent] = 0
		}
		dateBegin = timeBegin.Format(util.DateLayer)
	case "7d":
		timeBegin := nowTime.AddDate(0, 0, -7)
		for i := 0; i < 7; i++ {
			timeCurrent := timeBegin.AddDate(0, 0, i).Add(time.Duration(tz) * time.Hour).Format(util.DateLayer)
			amountStatSet[timeCurrent] = 0
		}
		dateBegin = timeBegin.Format(util.DateLayer)
	case "1m":
		timeBegin := nowTime.AddDate(0, -1, 0)
		for i := 0; i < 31; i++ {
			timeCurrent := timeBegin.AddDate(0, 0, i).Add(time.Duration(tz) * time.Hour)
			if timeCurrent.After(time.Now()) {
				break
			}
			amountStatSet[timeCurrent.Format(util.DateLayer)] = 0
		}
		dateBegin = timeBegin.Format(util.DateLayer)
	default:
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid chart gap")
	}

	// get payments
	paymentSet, paymentErr := db.FindApplicationFinishPaymentWithCurrencyByTime(application.ID, util.GetTimeStrFixedByTZ(dateBegin+" 00:00:00", tz), util.GetTimeStrFixedByTZ(dateToday+" 23:59:59", tz))
	if paymentErr != nil {
		return nil, paymentErr
	}

	for _, paymentItem := range paymentSet {
		if paymentItem.CollectionAmount.IsZero() {
			continue
		}

		canWithdrawAmount := paymentItem.CollectionAmount.Copy().Sub(paymentItem.PlatformFeeAmount).Sub(paymentItem.PoolFeeAmount).RawBigInt()
		paymentAmountRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(canWithdrawAmount, paymentItem.Currency)
		if exchangeErr != nil {
			continue
		}
		paymentAmount, _ := paymentAmountRaw.Float64()

		// timezone fixed
		paymentCreatedAtFixed := paymentItem.CreatedAt.Add(time.Duration(tz) * time.Hour)
		paymentTime := ""

		if gap == "24h" {
			paymentTime = paymentCreatedAtFixed.Format(util.TimeHourLayer)
		} else {
			paymentTime = paymentCreatedAtFixed.Format(util.DateLayer)
		}

		if _, ok2 := amountStatSet[paymentTime]; !ok2 {
			continue
		}

		amountStatSet[paymentTime] += paymentAmount
	}

	paymentAmountStat := make([]paymentAmountStatItem, 0)
	for title, amount := range amountStatSet {
		paymentAmountStat = append(paymentAmountStat, paymentAmountStatItem{
			Title:  title,
			Amount: amount,
		})
	}

	return &merchantPaymentStatChartResponse{
		Gap:               gap,
		PaymentAmountStat: paymentAmountStat,
	}, nil
}

func getMerchantDefaultApplication(merchantId, appId uint) (*db.Application, error) {
	var (
		application db.Application
		findAppErr  error
	)

	if appId == 0 {
		application, findAppErr = db.FindApplicationByMerchant(merchantId)
		if findAppErr != nil || application.ID == 0 {
			return nil, http_util.NewHttpError(http.StatusNotFound, "application not found")
		}
	} else {
		application, findAppErr = db.FindApplicationById(appId)
		if findAppErr != nil || application.ID == 0 {
			return nil, http_util.NewHttpError(http.StatusNotFound, "application not found")
		}

		if application.MerchantId != merchantId {
			return nil, http_util.NewHttpError(http.StatusForbidden, "application not belongs to you")
		}
	}

	return &application, nil
}
