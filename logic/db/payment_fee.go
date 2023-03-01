package db

import (
	"time"
	"veric-backend/logic/db/types"
)

const (
	PaymentFeeFeeTypePlatform = "Platform"
	PaymentFeeFeeTypePool     = "Pool"
)

type PaymentFee struct {
	CommonModel

	Date          time.Time     `gorm:"type:date;uniqueIndex:idx_pf_dtc,priority:1;not null" json:"date"`
	FeeType       string        `gorm:"uniqueIndex:idx_pf_dtc,priority:2;index:idx_pf_tc,priority:1;size:30;not null" json:"fee_type"`
	CurrencyId    uint          `gorm:"uniqueIndex:idx_pf_dtc,priority:3;index:idx_pf_tc,priority:2;not null" json:"currency_id"`
	Currency      *Currency     `gorm:"foreignKey:CurrencyId" json:"currency"`
	TotalAmount   *types.BigInt `gorm:"size:50;not null" json:"total_amount"`
	FrozenAmount  *types.BigInt `gorm:"size:50;not null" json:"frozen_amount"`
	ClaimedAmount *types.BigInt `gorm:"size:50;not null" json:"claimed_amount"`
}

func FindPaymentFeeWithCurrencyByType(feeType []string) (paymentFees []*PaymentFee, err error) {
	err = db.Preload("Currency").Where("fee_type IN ?", feeType).Find(&paymentFees).Error
	return
}

func FindPaymentFeeWithCurrencyByTypeDate(feeType, date string) (paymentFees []*PaymentFee, err error) {
	err = db.Preload("Currency").Where("fee_type = ? AND date = ?", feeType, date).Find(&paymentFees).Error
	return
}

func FindPaymentFeeByDateAndFeeTypeAndCurrencyId(feeType, date string, currencyId uint) (paymentFees *PaymentFee, err error) {
	err = db.Where("fee_type = ? AND date = ? AND currency_id = ?", feeType, date, currencyId).Find(&paymentFees).Error
	return
}

func FindAllAvailablePaymentFeesByFeeType(feeType []string, option ...Options) (paymentFees []*PaymentFee, err error) {
	err = useOptions(option...).Where("fee_type IN ? AND total_amount != claimed_amount", feeType).Find(&paymentFees).Error
	return
}

func FindAvailablePaymentFeesByFeeTypeAndCurrencyId(feeType string, currencyId uint) (paymentFees []*PaymentFee, err error) {
	err = db.
		Where("total_amount != claimed_amount and frozen_amount in ('', '0')").
		Where("fee_type = ? AND currency_id = ?", feeType, currencyId).
		Find(&paymentFees).
		Error
	return
}

func FindPaymentFeesById(ids []uint, option ...Options) (paymentFees []*PaymentFee, err error) {
	err = useOptions(option...).Where("id in ?", ids).Find(&paymentFees).Error
	return
}

func UpdatePaymentFeeTotalAmount(fee *PaymentFee, totalAmount *types.BigInt, option ...Options) error {
	tx := useOptions(option...).
		Model(&PaymentFee{}).
		Where("id = ? and total_amount = ?", fee.ID, fee.TotalAmount).
		Update("total_amount", totalAmount)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return ErrLockTryAgain
	} else {
		return nil
	}
}

func UpdatePaymentFeeFrozenAmount(fee *PaymentFee, frozenAmount *types.BigInt, option ...Options) error {
	tx := useOptions(option...).
		Model(&PaymentFee{}).
		Where("id = ? and frozen_amount = ?", fee.ID, fee.FrozenAmount).
		Update("frozen_amount", frozenAmount)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return ErrLockTryAgain
	} else {
		return nil
	}
}

func AddPaymentFeeClaimedAmountAndSubFrozenAmount(fee *PaymentFee, option ...Options) error {
	tx := useOptions(option...).
		Model(&PaymentFee{}).
		Where("id = ? and frozen_amount = ? and claimed_amount = ?", fee.ID, fee.FrozenAmount, fee.ClaimedAmount).
		UpdateColumns(map[string]any{
			"frozen_amount":  fee.FrozenAmount.Copy().SetZero(),
			"claimed_amount": fee.ClaimedAmount.Copy().Add(fee.FrozenAmount),
		})
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return ErrLockTryAgain
	} else {
		return nil
	}
}

func UpdatePaymentFeesFrozenAmountToZero(ids []uint, option ...Options) error {
	tx := useOptions(option...).
		Model(&PaymentFee{}).
		Where("id in ?", ids).
		UpdateColumns(map[string]any{
			"frozen_amount": types.NewBigIntFast(0),
		})
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return ErrLockTryAgain
	} else {
		return nil
	}
}

func SavePaymentFee(fee *PaymentFee) error {
	return db.Save(fee).Error
}
