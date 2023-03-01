package db

type Merchant struct {
	CommonModel

	Name         string        `gorm:"size:100;not null" json:"name"`
	Did          string        `gorm:"index:idx_merchant_did;size:100;not null" json:"did"`
	Applications []Application `json:"applications"`
	OwnerId      uint          `gorm:"index:idx_merchant_owner_id;not null" json:"owner_id"`
	Status       string        `gorm:"index:idx_merchant_status;size:100;not null" json:"status"`
}

func FindMerchantById(id uint) (merchant *Merchant, err error) {
	err = db.Where("id = ?", id).Find(&merchant).Error
	return
}

func FindMerchantByOwnerId(ownerId uint) (merchant *Merchant, err error) {
	err = db.Where("owner_id = ?", ownerId).Find(&merchant).Error
	return
}

func GetMerchantCountByStatus(status string) uint {
	var count int64

	query := db.Table("merchants")

	if status != "all" {
		query.Where("status = ?", status)
	}

	query.Count(&count)

	return uint(count)
}

func GetMerchantCountByStatusCreatedAt(status, start, end string) uint {
	var count int64

	query := db.Table("merchants")

	if status != "all" {
		query.Where("status = ?", status)
	}

	query.Where("created_at BETWEEN ? AND ?", start, end).Count(&count)

	return uint(count)
}

func FindMerchantWithApplicationsByDid(did string) (merchant *Merchant, err error) {
	err = db.Preload("Applications").Where("did = ?", did).Find(&merchant).Error
	return
}

func FindPaginationMerchantByStatus(status string, offset int, size int) (merchants []*Merchant, err error) {
	query := db.Preload("Applications").Table("merchants")

	if status != "all" {
		query.Where("status = ?", status)
	}

	err = query.Limit(size).Offset(offset).Order("id desc").Find(&merchants).Error
	return
}

func SaveMerchant(merchant *Merchant) (err error) {
	return db.Save(merchant).Error
}

func CreateMerchant(merchant *Merchant) (err error) {
	return db.Create(merchant).Error
}
