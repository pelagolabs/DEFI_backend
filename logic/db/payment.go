package db

import (
	"time"
	"veric-backend/logic/db/types"
)

const (
	PaymentStatusCreated = "created"
	PaymentStatusPending = "pending"
	PaymentStatusSuccess = "success"
	PaymentStatusClosed  = "closed"
)

const (
	AmountStatusOverpaid  = "overpaid"
	AmountStatusUnderpaid = "underpaid"
	AmountStatusUnpaid    = "unpaid"
	AmountStatusPaid      = "paid"
)

type Payment struct {
	CommonModel

	MerchantId uint      `gorm:"index:idx_merchant_id;not null" json:"merchant_id"`
	Merchant   *Merchant `gorm:"foreignKey:MerchantId" json:"merchant"`

	ApplicationId uint         `gorm:"index:idx_application_id;not null" json:"application_id"`
	Application   *Application `gorm:"foreignKey:ApplicationId" json:"application"`

	CurrencyId uint      `gorm:"index:idx_currency_id;not null" json:"currency_id"`
	Currency   *Currency `gorm:"foreignKey:CurrencyId" json:"currency"`

	PaymentTxs []*PaymentTx `json:"payment_txs"`
	VCs        []*VC        `json:"vcs"`

	CurrencyPriceInUSD     float64       `gorm:"type:decimal(40,20);not null" json:"currency_price_in_usd"`
	PaymentNum             string        `gorm:"index:idx_payment_num;size:100;not null" json:"payment_num"`
	Amount                 *types.BigInt `gorm:"size:50;not null" json:"amount"`
	AmountInCent           uint64        `gorm:"not null;default:0" json:"amount_in_cent"`
	CollectionAmount       *types.BigInt `gorm:"size:50;not null" json:"collection_amount"`
	CollectionAmountInCent uint64        `gorm:"not null;default:0" json:"collection_amount_in_cent"`
	CollectionAddress      string        `gorm:"index:idx_collection_address;size:50;not null" json:"collection_address"`
	Status                 string        `gorm:"index:idx_collection_address;index:idx_payment_status;size:100;not null" json:"status"`
	AmountStatus           string        `gorm:"size:20;not null" json:"amount_status"`
	Title                  string        `gorm:"size:50;not null" json:"title"`
	Desc                   string        `gorm:"size:50;not null" json:"desc"`
	FinishTime             time.Time     `gorm:"index:idx_payment_finish_time;not null" json:"finish_time"`
	DeadlineTime           time.Time     `gorm:"not null" json:"deadline_time"`
	PlatformFeeAmount      *types.BigInt `gorm:"size:50;not null" json:"platform_fee_amount"`
	PoolFeeAmount          *types.BigInt `gorm:"size:50;not null" json:"pool_fee_amount"`
	Slippage               float64       `gorm:"not null;default:1" json:"slippage"`
}

type PaymentCollectionAmount struct {
	CurrencyId           uint   `json:"currency_id"`
	CurrencySearchName   string `json:"currency_search_name"`
	CurrencyDecimalCount uint   `json:"currency_decimal_count"`
	CollectionAmount     string `json:"collection_amount"`
}

func FindPaymentById(id uint) (payment *Payment, err error) {
	err = db.
		Preload("Currency.Blockchain").
		Preload("Application").
		Where("id = ?", id).
		Find(&payment).
		Error
	return
}

func FindPaymentByPaymentNum(paymentNum string) (payment *Payment, err error) {
	err = db.Where("payment_num = ?", paymentNum).Find(&payment).Error
	return
}

func FindPaymentsByStatus(status []string) (payments []*Payment, err error) {
	err = db.Where("status IN ?", status).Find(&payments).Error
	return
}

func FindPaymentByCollectionAddress(collectionAddress string) (payment *Payment, err error) {
	err = db.Preload("Currency").
		Where("collection_address = ? and status in ?", collectionAddress, []string{PaymentStatusCreated, PaymentStatusPending}).
		Find(&payment).
		Error
	return
}

func FindFullPaymentByPaymentNum(paymentNum string) (payment *Payment, err error) {
	err = db.
		Preload("Merchant").
		Preload("Application.Currencies").
		Preload("Currency").
		Where("payment_num = ?", paymentNum).
		Find(&payment).
		Error
	return
}

func FindDeadPayment() (payment []*Payment, err error) {
	err = db.Preload("Currency").
		Where("deadline_time < ? and status in ?", time.Now(), []string{PaymentStatusCreated, PaymentStatusPending}).
		Find(&payment).
		Order("id asc").
		Limit(100).
		Error
	return
}

func FindApplicationPaymentWithCurrencyVCByPaymentNum(applicationId uint, paymentNum string) (payment Payment, err error) {
	err = db.Preload("Currency").Preload("VCs").Where("application_id = ? AND payment_num = ? AND collection_amount != '0'", applicationId, paymentNum).Find(&payment).Error
	return
}

func FindApplicationPaymentById(applicationId, id uint) (payment Payment, err error) {
	err = db.Where("application_id = ? AND id = ?", applicationId, id).Find(&payment).Error
	return
}

func FindApplicationPaymentWithCurrencyTxById(applicationId, id uint) (payment Payment, err error) {
	err = db.Preload("PaymentTxs").Preload("Currency").Where("application_id = ? AND id = ?", applicationId, id).Find(&payment).Error
	return
}

func FindApplicationPaymentWithCurrencyTxByPaymentNum(applicationId uint, paymentNum string) (payment Payment, err error) {
	err = db.Preload("PaymentTxs").Preload("Currency").Where("application_id = ? AND payment_num = ?", applicationId, paymentNum).Find(&payment).Error
	return
}

func GetApplicationPaymentCountByStatusCurrencyDate(applicationId uint, status string, currencyId uint, date string, displayAll int) uint {
	var count int64

	query := db.Table("payments")

	if status == "normal" {
		query.Where("status IN ?", []string{PaymentStatusPending, PaymentStatusSuccess, PaymentStatusClosed})
	} else if status != "" {
		query.Where("status = ?", status)
	}

	if currencyId > 0 {
		query.Where("currency_id = ?", currencyId)
	}

	if date != "" {
		query.Where("created_at BETWEEN ? AND ?", date+" 00:00:00", date+" 23:59:59")
	}

	if displayAll == 0 {
		query.Where("collection_amount != '0'")
	}

	query.Where("application_id = ?", applicationId).Count(&count)

	return uint(count)
}

func FindPaginationApplicationPaymentWithCurrencyVCByStatusCurrencyDate(applicationId uint, status string, currencyId uint, date string, offset int, size int, displayAll int) (payments []Payment, err error) {
	query := db.Table("payments")

	if status == "normal" {
		query.Where("status IN ?", []string{PaymentStatusPending, PaymentStatusSuccess, PaymentStatusClosed})
	} else if status != "" {
		query.Where("status = ?", status)
	}

	if currencyId > 0 {
		query.Where("currency_id = ?", currencyId)
	}

	if date != "" {
		query.Where("created_at BETWEEN ? AND ?", date+" 00:00:00", date+" 23:59:59")
	}

	if displayAll == 0 {
		query.Where("collection_amount != '0'")
	}

	err = query.Preload("Currency").Preload("VCs").Where("application_id = ?", applicationId).Limit(size).Offset(offset).Order("id desc").Find(&payments).Error
	return
}

func FindApplicationFinishPaymentWithCurrencyByTime(applicationId uint, start, end string) (payments []Payment, err error) {
	query := db.Table("payments")

	if start != "all" {
		query.Where("payments.created_at BETWEEN ? AND ?", start, end)
	}

	err = query.Preload("Currency").Where("application_id = ? AND payments.status IN ?", applicationId, []string{PaymentStatusSuccess, PaymentStatusClosed}).Find(&payments).Error

	return
}

func SavePayment(payment *Payment) (err error) {
	return db.Save(payment).Error
}

func SavePaymentWithStatusLock(payment *Payment, oriStatus string, options ...Options) (err error) {
	tx := useOptions(options...).Model(&Payment{}).Where("id = ? and status = ?", payment.ID, oriStatus).Updates(payment)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return ErrLockTryAgain
	} else {
		return nil
	}
}
