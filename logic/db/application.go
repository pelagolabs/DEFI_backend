package db

import (
	"time"
)

type Application struct {
	CommonModel

	MerchantId      uint       `gorm:"index:idx_application_merchant_id;not null" json:"merchant_id"`
	Name            string     `gorm:"size:100;not null" json:"name"`
	Describe        string     `gorm:"size:200" json:"describe"`
	Link            string     `gorm:"size:200" json:"link"`
	Slippage        float64    `gorm:"not null;default:1" json:"slippage"`
	Currencies      []Currency `gorm:"many2many:application_currencies;" json:"currencies"`
	CallbackUrl     string     `gorm:"size:200" json:"callbackUrl"`
	ApiKey          string     `gorm:"size:100;not null" json:"api_key"`
	ApiKeyCreatedAt time.Time  `gorm:"not null" json:"api_key_created_at"`
	IpnKey          string     `gorm:"size:100;not null" json:"ipn_key"`
	IpnKeyCreatedAt time.Time  `gorm:"not null" json:"ipn_key_created_at"`
	Status          string     `gorm:"index:idx_application_status;size:100;not null" json:"status"`
	LegalTender     string     `gorm:"size:30;not null;default:'usd'" json:"legal_tender"`
}

func FindApplicationById(id uint) (application Application, err error) {
	err = db.Where("id = ?", id).Find(&application).Error
	return
}

func FindApplicationByMerchant(merchantId uint) (application Application, err error) {
	err = db.Where("merchant_id = ?", merchantId).Find(&application).Error
	return
}

func FindApplicationByApiKey(apiKey string) (application Application, err error) {
	err = db.Where("api_Key = ?", apiKey).Find(&application).Error
	return
}

func ReplaceApplicationCurrency(application *Application, currencyIds []uint) (err error) {
	newCurrencySet := make([]Currency, 0)
	for _, currencyId := range currencyIds {
		newCurrencySet = append(newCurrencySet, Currency{
			CommonModel: CommonModel{
				ID: currencyId,
			},
		})
	}

	err = db.Model(application).Association("Currencies").Replace(newCurrencySet)
	return
}

func GetApplicationCountByStatusMerchant(status string, merchantId uint) uint {
	var count int64

	query := db.Table("applications")

	if status != "all" {
		query.Where("status = ?", status)
	}

	if merchantId != 0 {
		query.Where("merchant_id = ?", merchantId)
	}

	query.Count(&count)

	return uint(count)
}

func FindPaginationApplicationByStatusMerchant(status string, offset int, size int, merchantId uint) (applications []Application, err error) {
	query := db.Preload("Currencies").Table("applications")

	if status != "all" {
		query.Where("status = ?", status)
	}

	if merchantId != 0 {
		query.Where("merchant_id = ?", merchantId)
	}

	err = query.Limit(size).Offset(offset).Order("id desc").Find(&applications).Error
	return
}

func SaveApplication(application *Application) (err error) {
	return db.Save(application).Error
}
