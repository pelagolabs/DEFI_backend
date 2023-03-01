package db

import "veric-backend/logic/db/types"

type FeeWithdrawLog struct {
	CommonModel

	FeeType    string        `gorm:"index:idx_fee_withdraw_log_type;size:30;not null" json:"fee_type"`
	CurrencyId uint          `gorm:"index:idx_fee_withdraw_log_currency_id;not null" json:"currency_id"`
	Currency   *Currency     `gorm:"foreignKey:CurrencyId" json:"currency"`
	ToAddress  string        `gorm:"size:50;not null" json:"to_address"`
	Amount     *types.BigInt `gorm:"size:50;not null" json:"amount"`
	ManagerId  uint          `gorm:"not null" json:"manager_id"`
	Remark     string        `gorm:"size:500;default:''" json:"remark"`
	TxHash     string        `gorm:"size:100;not null;default:''" json:"tx_hash"`
}

func GetFeeWithdrawLogCountByStatusCurrencyDate(currencyId uint, date string) uint {
	var count int64

	query := db.Table("fee_withdraw_logs")

	if currencyId > 0 {
		query.Where("currency_id = ?", currencyId)
	}

	if date != "" {
		query.Where("created_at BETWEEN ? AND ?", date+" 00:00:00", date+" 23:59:59")
	}

	query.Count(&count)

	return uint(count)
}

func FindPaginationFeeWithdrawLogByStatusCurrencyDate(currencyId uint, date string, offset int, size int) (withdrawLogs []FeeWithdrawLog, err error) {
	query := db.Table("fee_withdraw_logs")

	if currencyId > 0 {
		query.Where("currency_id = ?", currencyId)
	}

	if date != "" {
		query.Where("created_at BETWEEN ? AND ?", date+" 00:00:00", date+" 23:59:59")
	}

	err = query.Preload("Currency").Preload("Currency.Blockchain").Limit(size).Offset(offset).Order("id desc").Find(&withdrawLogs).Error
	return
}

func SaveFeeWithdrawLog(log *FeeWithdrawLog) (err error) {
	return db.Save(log).Error
}
