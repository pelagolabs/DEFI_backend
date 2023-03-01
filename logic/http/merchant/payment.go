package merchant

import (
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
	"veric-backend/internal/log"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/blockchain/exchange"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/http/http_util"
)

func GetPaymentSummary(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	params := mux.Vars(r.Request)

	tz := r.QueryWithDefaultInt("tz", 0)

	// check app id
	if _, ok := params["app_id"]; !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	appId, atoiErr := strconv.Atoi(params["app_id"])
	if atoiErr != nil {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	// get user jwt claim
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, uint(appId))
	if err != nil {
		return nil, err
	}

	// fmt.Printf("now: %s fixed-now: %s\n", time.Now(), time.Now().Add(time.Duration(tz)*time.Hour))

	nowTime := time.Now().Add(time.Duration(tz) * time.Hour)
	dateBegin := nowTime.AddDate(0, 0, -90).Format(util.DateLayer)
	dateToday := nowTime.Format(util.DateLayer)
	paymentSummary := &paymentSummaryResponse{
		Latest90DaysTotalPayment: 0,
		TodayTotalPayment:        0,
	}

	// get latest 90 days summary
	latest90Stats, sum90Err := db.FindApplicationFinishPaymentWithCurrencyByTime(application.ID, util.GetTimeStrFixedByTZ(dateBegin+" 00:00:00", tz), util.GetTimeStrFixedByTZ(dateToday+" 23:59:59", tz))
	if sum90Err != nil {
		return nil, sum90Err
	}

	for _, paymentItem := range latest90Stats {
		if paymentItem.CollectionAmount.IsZero() {
			continue
		}

		canWithdrawAmount := paymentItem.CollectionAmount.Copy().Sub(paymentItem.PlatformFeeAmount).Sub(paymentItem.PoolFeeAmount).RawBigInt()
		paymentAmountRaw, exchangeErr := exchange.DefaultManage.ExchangeCoinToUSD(canWithdrawAmount, paymentItem.Currency)
		if exchangeErr != nil {
			continue
		}

		paymentAmount, _ := paymentAmountRaw.Float64()
		paymentSummary.Latest90DaysTotalPayment += paymentAmount
	}

	// fmt.Printf("start: %s end: %s\n", util.GetTimeStrFixedByTZ(dateToday+" 00:00:00", tz), util.GetTimeStrFixedByTZ(dateToday+" 23:59:59", tz))

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
		paymentSummary.TodayTotalPayment += paymentAmount
	}

	return paymentSummary, nil
}

func GetPaymentList(typ interface{}, r *http_util.HTTPContext) (resp interface{}, respErr error) {
	status := r.QueryWithDefault("status", "normal")
	page := r.QueryWithDefaultInt("page", 1)
	size := r.QueryWithDefaultInt("size", 10)
	paymentNum := r.QueryWithDefault("payment_num", "")
	currencyId := uint(r.QueryWithDefaultInt("currency_id", 0))
	date := r.QueryWithDefault("date", "")
	displayAll := r.QueryWithDefaultInt("display_all", 0)

	// check app id
	params := mux.Vars(r.Request)
	if _, ok := params["app_id"]; !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	appId, atoiErr := strconv.Atoi(params["app_id"])
	if atoiErr != nil {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager", "employee"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, uint(appId))
	if err != nil {
		return nil, err
	}

	// calc pagination
	offset := (page - 1) * size

	// get transaction list
	var (
		payments = make([]paymentListItem, 0)
		total    uint
	)

	if paymentNum != "" {
		payment, findErr := db.FindApplicationPaymentWithCurrencyVCByPaymentNum(application.ID, paymentNum)
		if findErr != nil {
			log.GetLogger().Error("get payment list failed", zap.String("error", findErr.Error()))
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "get payment list failed")
		}

		if payment.ID > 0 {
			total = 1

			canWithdrawAmount := payment.CollectionAmount.Copy().Sub(payment.PlatformFeeAmount).Sub(payment.PoolFeeAmount).RawBigInt()
			amount, _ := exchange.DefaultManage.ExchangeCoinToNaturalAmount(canWithdrawAmount, payment.Currency).Float64()

			vcs := make([]paymentVCItem, 0)
			for _, vc := range payment.VCs {
				vcs = append(vcs, paymentVCItem{
					VCID:     vc.VCID,
					VCStatus: vc.VCStatus,
				})
			}

			payments = append(payments, paymentListItem{
				PaymentNum:       payment.PaymentNum,
				CurrencySymbol:   payment.Currency.Symbol,
				Amount:           amount,
				Status:           payment.Status,
				VCs:              vcs,
				CreatedAt:        payment.CreatedAt,
				FinishTime:       payment.FinishTime,
				OrderAmount:      payment.Amount.String(),
				CollectionAmount: payment.CollectionAmount.String(),
				Slippage:         payment.Slippage,
			})
		}
	} else {
		total = db.GetApplicationPaymentCountByStatusCurrencyDate(application.ID, status, currencyId, date, displayAll)
		paymentSet, findErr := db.FindPaginationApplicationPaymentWithCurrencyVCByStatusCurrencyDate(application.ID, status, currencyId, date, offset, size, displayAll)
		if findErr != nil {
			log.GetLogger().Error("get payment list failed", zap.String("error", findErr.Error()))
			return nil, http_util.NewHttpError(http.StatusInternalServerError, "get payment list failed")
		}

		for _, payment := range paymentSet {
			canWithdrawAmount := payment.CollectionAmount.Copy().Sub(payment.PlatformFeeAmount).Sub(payment.PoolFeeAmount).RawBigInt()
			amount, _ := exchange.DefaultManage.ExchangeCoinToNaturalAmount(canWithdrawAmount, payment.Currency).Float64()

			vcs := make([]paymentVCItem, 0)
			for _, vc := range payment.VCs {
				vcs = append(vcs, paymentVCItem{
					VCID:     vc.VCID,
					VCStatus: vc.VCStatus,
				})
			}

			payments = append(payments, paymentListItem{
				PaymentNum:       payment.PaymentNum,
				CurrencySymbol:   payment.Currency.Symbol,
				Amount:           amount,
				Status:           payment.Status,
				VCs:              vcs,
				CreatedAt:        payment.CreatedAt,
				FinishTime:       payment.FinishTime,
				OrderAmount:      payment.Amount.String(),
				CollectionAmount: payment.CollectionAmount.String(),
				Slippage:         payment.Slippage,
			})
		}
	}

	return &paymentListResponse{
		Total:    total,
		Payments: payments,
	}, nil
}

func GetPaymentDetail(r *http_util.HTTPContext) (resp interface{}, respErr error) {
	params := mux.Vars(r.Request)

	// check app id
	if _, ok := params["app_id"]; !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	appId, atoiErr := strconv.Atoi(params["app_id"])
	if atoiErr != nil {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	// check payment id
	if _, ok := params["payment_num"]; !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid payment num")
	}
	paymentNum := params["payment_num"]

	// get user
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager", "employee"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, uint(appId))
	if err != nil {
		return nil, err
	}

	payment, paymentErr := db.FindApplicationPaymentWithCurrencyTxByPaymentNum(application.ID, paymentNum)
	if paymentErr != nil {
		return nil, paymentErr
	}

	if payment.ID == 0 {
		return nil, http_util.NewHttpError(http.StatusNotFound, "payment not found")
	}

	var (
		completedAt  = "-"
		paymentHash  = make([]string, 0)
		payInAddress = "-"
		amount       = float64(0)
	)

	switch payment.Status {
	case db.PaymentStatusPending:
		for _, tx := range payment.PaymentTxs {
			txHashSet := strings.Split(tx.Hash, ":")
			paymentHash = append(paymentHash, txHashSet[0])
			payInAddress = tx.TxAddress
		}
	case db.PaymentStatusSuccess:
		for _, tx := range payment.PaymentTxs {
			txHashSet := strings.Split(tx.Hash, ":")
			paymentHash = append(paymentHash, txHashSet[0])
			payInAddress = tx.TxAddress
		}
		completedAt = payment.UpdatedAt.Format(util.TimeLayer)
	case db.PaymentStatusClosed:
		if !payment.CollectionAmount.IsZero() {
			for _, tx := range payment.PaymentTxs {
				txHashSet := strings.Split(tx.Hash, ":")
				paymentHash = append(paymentHash, txHashSet[0])
				payInAddress = tx.TxAddress
			}
		}
		completedAt = payment.UpdatedAt.Format(util.TimeLayer)
	}

	amount, _ = exchange.DefaultManage.ExchangeCoinToNaturalAmount(payment.Amount.RawBigInt(), payment.Currency).Float64()

	paymentDetail := &paymentDetailResponse{
		PaymentNum:        payment.PaymentNum,
		OriginAmount:      float64(payment.AmountInCent) / 100,
		PayAmount:         amount,
		PayCurrencySymbol: payment.Currency.Symbol,
		ExchangeRate:      float64(payment.CollectionAmountInCent) / float64(payment.AmountInCent),
		OutcomeAmount:     float64(payment.CollectionAmountInCent) / 100,
		PayInAddress:      payInAddress,
		PayOutAddress:     payment.CollectionAddress,
		PaymentHash:       paymentHash,
		Status:            payment.Status,
		CreatedAt:         payment.CreatedAt.Format(util.TimeLayer),
		CompletedAt:       completedAt,
	}

	return paymentDetail, nil
}

func NewPayment(req *newPaymentRequest, r *http_util.HTTPContext) (resp interface{}, err error) {
	params := mux.Vars(r.Request)

	// check app id
	if _, ok := params["app_id"]; !ok {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	appId, atoiErr := strconv.Atoi(params["app_id"])
	if atoiErr != nil {
		return nil, http_util.NewHttpError(http.StatusBadRequest, "invalid application id")
	}

	// get user jwt claim
	jwtClaims, jwtErr := r.GetMerchantJwt("user-auth-token")
	if jwtErr != nil {
		return nil, http_util.NewHttpError(http.StatusUnauthorized, "please login first")
	}

	if !r.CheckPermissionAny([]string{jwtClaims.Role}, []string{"admin", "shop_manager", "employee"}) {
		return nil, http_util.NewHttpError(http.StatusForbidden, "no permission for this operation")
	}

	// check global merchant status
	if config.Get().Logic.SuspendAllMerchant {
		return nil, errors.New("system operate paused. please connect to admin")
	}

	// get application
	application, err := getMerchantDefaultApplication(jwtClaims.MerchantId, uint(appId))
	if err != nil {
		return nil, err
	}

	// check merchant status
	merchant, err := db.FindMerchantById(application.MerchantId)
	if err != nil {
		return nil, err
	}

	if merchant.Status == "unavailable" {
		return nil, errors.New("merchant was suspend. please connect to admin")
	}

	currency, err := db.FindCurrencyByChainAndSymbol(req.ChainId, req.Currency, db.WithPreload("Blockchain"))
	if err != nil {
		return nil, err
	}

	if application.ID == 0 || currency.ID == 0 {
		return nil, errors.New("AccessKey, ChainId or Currency is invalid")
	}

	if currency.Blockchain.Status == "unavailable" {
		return nil, errors.New("chain was suspend. please connect to admin")
	}

	if currency.Status == "unavailable" {
		return nil, errors.New("currency was suspend. please connect to admin")
	}

	price, err := exchange.DefaultManage.QueryPrice(currency)
	if err != nil {
		return nil, err
	}

	coin, err := exchange.DefaultManage.ExchangeUSDToCoin(big.NewFloat(float64(req.AmountInCent)/100), currency)
	if err != nil {
		return nil, err
	}

	payment := &db.Payment{
		PaymentNum:    util.RandString(16),
		AmountInCent:  req.AmountInCent,
		MerchantId:    application.MerchantId,
		ApplicationId: application.ID,
		Status:        db.PaymentStatusCreated,
		AmountStatus:  db.AmountStatusUnpaid,
		Title:         "New Payment",
		Desc:          "the payment created in merchant dashboard",
		CurrencyId:    currency.ID,
		Amount:        types.NewBigInt(coin),
		DeadlineTime:  time.Now().Add(90 * time.Minute),
		FinishTime:    time.Now().Add(30 * time.Minute),
		Slippage:      application.Slippage,
	}

	payment.CurrencyPriceInUSD, _ = price.Float64()
	payment.CollectionAddress, err = contract.DefaultContract.GetAddressFromPool(currency)
	if err != nil {
		return nil, err
	}

	err = db.SavePayment(payment)
	if err != nil {
		return nil, err
	}

	return &newPaymentResponse{
		CurrencyName:         currency.Name,
		CurrencyDecimalCount: currency.DecimalCount,
		PaymentNum:           payment.PaymentNum,
		Amount:               payment.Amount.String(),
		FriendlyAmount:       payment.Amount.Copy().Ceil(int(currency.FriendlyDecimalCount)).String(),
		AmountInCent:         payment.AmountInCent,
		CollectionAddress:    payment.CollectionAddress,
		Status:               payment.Status,
		Title:                payment.Title,
		Desc:                 payment.Desc,
		FinishTime:           payment.FinishTime,
	}, nil
}
