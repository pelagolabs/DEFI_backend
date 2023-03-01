package db

const (
	WithdrawStatusCreated  = "created"
	WithdrawStatusComplete = "complete"
	WithdrawStatusClosed   = "closed"
)

type Withdraw struct {
	CommonModel

	MerchantId uint      `gorm:"index:idx_withdraw_merchant_id;not null" json:"merchant_id"`
	Merchant   *Merchant `gorm:"foreignKey:MerchantId" json:"merchant"`

	ApplicationId uint         `gorm:"index:idx_withdraw_application_id;not null" json:"application_id"`
	Application   *Application `gorm:"foreignKey:ApplicationId" json:"application"`

	CurrencyId   uint      `gorm:"index:idx_withdraw_currency_id;not null" json:"currency_id"`
	Currency     *Currency `gorm:"foreignKey:CurrencyId" json:"currency"`
	WithdrawNum  string    `gorm:"index:idx_withdraw_num;size:100;not null" json:"withdraw_num"`
	Amount       string    `gorm:"size:50;not null" json:"amount"`
	AmountInCent uint64    `gorm:"not null;default:0" json:"amount_in_cent"`

	Status string `gorm:"index:idx_withdraw_status;size:100;not null" json:"status"`
	TxHash string `gorm:"size:100;not null;default:''" json:"tx_hash"`
}

type WithdrawStat struct {
	CurrencyId           uint    `json:"currency_id"`
	CurrencyPrice        float64 `json:"currency_price"`
	CurrencyDecimalCount uint    `json:"currency_decimal_count"`
	TotalAmount          float64 `json:"total_amount"`
}

func FindWithdrawById(id uint) (withdraw *Withdraw, err error) {
	err = db.Where("id = ?", id).Find(&withdraw).Error
	return
}

func FindApplicationWithdrawById(applicationId, id uint) (withdraw *Withdraw, err error) {
	err = db.Where("application_id = ? AND id = ?", applicationId, id).Find(&withdraw).Error
	return
}

func FindApplicationWithdrawWithCurrencyByPaymentNum(applicationId uint, withdrawNum string) (withdraw Withdraw, err error) {
	err = db.Preload("Currency").Where("application_id = ? AND withdraw_num = ?", applicationId, withdrawNum).Find(&withdraw).Error
	return
}

func GetApplicationWithdrawCountByStatusCurrencyDate(applicationId uint, status string, currencyId uint, date string) uint {
	var count int64

	query := db.Table("withdraws")

	if status != "all" {
		query.Where("status = ?", status)
	}

	if currencyId > 0 {
		query.Where("currency_id = ?", currencyId)
	}

	if date != "" {
		query.Where("created_at BETWEEN ? AND ?", date+" 00:00:00", date+" 23:59:59")
	}

	query.Where("application_id = ?", applicationId).Count(&count)

	return uint(count)
}

func FindPaginationApplicationWithdrawByStatusCurrencyDate(applicationId uint, status string, currencyId uint, date string, offset int, size int) (withdraws []Withdraw, err error) {
	query := db.Table("withdraws")

	if status != "all" {
		query.Where("status = ?", status)
	}

	if currencyId > 0 {
		query.Where("currency_id = ?", currencyId)
	}

	if date != "" {
		query.Where("created_at BETWEEN ? AND ?", date+" 00:00:00", date+" 23:59:59")
	}

	err = query.Preload("Currency").Where("application_id = ?", applicationId).Limit(size).Offset(offset).Order("id desc").Find(&withdraws).Error
	return
}

func FindApplicationFinishWithdrawWithCurrencyByTime(applicationId uint, start, end string) (withdraw []Withdraw, err error) {
	query := db.Table("withdraws")

	if start != "all" {
		query.Where("withdraws.created_at BETWEEN ? AND ?", start, end)
	}

	err = query.Preload("Currency").Where("application_id = ? AND withdraws.status IN ?", applicationId, []string{WithdrawStatusComplete}).Find(&withdraw).Error

	return
}

func SaveWithdraw(w *Withdraw, option ...Options) error {
	return useOptions(option...).Save(w).Error
}
