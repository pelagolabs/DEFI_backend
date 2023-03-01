package db

import (
	"time"
	"veric-backend/logic/db/types"
)

type PaymentBalance struct {
	CommonModel

	Date          time.Time     `gorm:"type:date;index:idx_pb_adc,priority:2;not null" json:"date"`
	ApplicationId uint          `gorm:"index:idx_pb_adc,priority:1;not null" json:"application_id"`
	Application   *Application  `gorm:"foreignKey:ApplicationId" json:"application"`
	CurrencyId    uint          `gorm:"index:idx_pb_adc,priority:3;not null" json:"currency_id"`
	Currency      *Currency     `gorm:"foreignKey:CurrencyId" json:"currency"`
	TotalAmount   *types.BigInt `gorm:"size:50;not null" json:"total_amount"`
	ClaimedAmount *types.BigInt `gorm:"size:50;not null" json:"claimed_amount"`
}

func FindPaymentBalanceWithCurrencyByApp(appId uint) (paymentBalances []*PaymentBalance, err error) {
	err = db.Preload("Currency").Where("application_id = ?", appId).Find(&paymentBalances).Error
	return
}

func FindUnClaimedPaymentBalanceWithCurrencyByApp(appId uint) (paymentBalances []*PaymentBalance, err error) {
	err = db.Preload("Currency").Where("application_id = ? AND total_amount != claimed_amount", appId).Find(&paymentBalances).Error
	return
}

func FindAllUnClaimedPaymentBalance(options ...Options) (paymentBalances []*PaymentBalance, err error) {
	err = useOptions(options...).Where("total_amount != claimed_amount").Find(&paymentBalances).Error
	return
}

func FindPaymentBalanceByApplicationIdAndDateAndCurrencyId(applicationId uint, date string, currencyId uint) (paymentBalance *PaymentBalance, err error) {
	err = db.Where("application_id = ? AND date = ? AND currency_id = ?", applicationId, date, currencyId).Find(&paymentBalance).Error
	return
}

func SavePaymentBalance(balance *PaymentBalance) error {
	return db.Save(balance).Error
}

func UpdatePaymentBalanceTotalAmount(balance *PaymentBalance, totalAmount *types.BigInt, option ...Options) error {
	tx := useOptions(option...).
		Model(&PaymentBalance{}).
		Where("id = ? and total_amount = ?", balance.ID, balance.TotalAmount).
		Update("total_amount", totalAmount)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return ErrLockTryAgain
	} else {
		return nil
	}
}

func UpdatePaymentBalanceClaimedAmount(balance *PaymentBalance, claimedAmount *types.BigInt, option ...Options) error {
	tx := useOptions(option...).
		Model(&PaymentBalance{}).
		Where("id = ? and claimed_amount = ?", balance.ID, balance.ClaimedAmount).
		Update("claimed_amount", claimedAmount)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return ErrLockTryAgain
	} else {
		return nil
	}
}
