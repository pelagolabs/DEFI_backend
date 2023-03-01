package open_api

import (
	"errors"
	"math/big"
	"time"
	"veric-backend/internal/util"
	"veric-backend/logic/blockchain/contract"
	"veric-backend/logic/blockchain/exchange"
	"veric-backend/logic/config"
	"veric-backend/logic/db"
	"veric-backend/logic/db/types"
	"veric-backend/logic/http/http_util"
)

type newPaymentRequest struct {
	AccessKey    string `json:"access_key" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Desc         string `json:"desc" validate:"required"`
	AmountInCent uint64 `json:"amount_in_cent" validate:"required"`
	ChainId      uint   `json:"chain_id" validate:"required"`
	Currency     string `json:"currency" validate:"required"`
}

type newPaymentResponse struct {
	CurrencyName         string    `json:"currency_name"`
	CurrencyDecimalCount uint      `json:"currency_decimal_count"`
	PaymentNum           string    `json:"payment_num"`
	Amount               string    `json:"amount"`
	FriendlyAmount       string    `json:"friendly_amount"`
	AmountInCent         uint64    `json:"amount_in_cent"`
	CollectionAddress    string    `json:"collection_address"`
	Status               string    `json:"status"`
	Title                string    `json:"title"`
	Desc                 string    `json:"desc"`
	FinishTime           time.Time `json:"finish_time"`
}

func NewPayment(req *newPaymentRequest, r *http_util.HTTPContext) (resp interface{}, err error) {
	application, err := db.FindApplicationByApiKey(req.AccessKey)
	if err != nil {
		return nil, err
	}

	// check global merchant status
	if config.Get().Logic.SuspendAllMerchant {
		return nil, errors.New("system operate paused. please connect to admin")
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
		Title:         req.Title,
		Desc:          req.Desc,
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

func FindFullPayment(r *http_util.HTTPContext) (resp interface{}, err error) {
	paymentNum := r.URL.Query().Get("payment_num")
	payment, err := db.FindFullPaymentByPaymentNum(paymentNum)
	if err != nil {
		return nil, err
	}

	if payment.ID == 0 {
		return nil, nil
	}

	return &newPaymentResponse{
		CurrencyName:         payment.Currency.Name,
		CurrencyDecimalCount: payment.Currency.DecimalCount,
		PaymentNum:           payment.PaymentNum,
		Amount:               payment.Amount.String(),
		FriendlyAmount:       payment.Amount.Copy().Ceil(int(payment.Currency.FriendlyDecimalCount)).String(),
		AmountInCent:         payment.AmountInCent,
		CollectionAddress:    payment.CollectionAddress,
		Status:               payment.Status,
		Title:                payment.Title,
		Desc:                 payment.Desc,
		FinishTime:           payment.FinishTime,
	}, nil
}
