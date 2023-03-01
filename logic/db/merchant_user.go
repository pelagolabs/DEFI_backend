package db

type MerchantUser struct {
	CommonModel

	MerchantId uint     `gorm:"index:idx_mu_merchant_id;not null" json:"merchant_id"`
	Merchant   Merchant `json:"merchant_detail"`
	UserId     uint     `gorm:"index:idx_mu_user_id;not null" json:"user_id"`
	User       User     `json:"user_detail"`
	Role       string   `gorm:"size:50;not null" json:"role"`
	Status     string   `gorm:"index:idx_user_status;size:50;not null" json:"status"`
}

func FindMerchantUserById(id uint) (merchantUser MerchantUser, err error) {
	err = db.Where("id = ?", id).Find(&merchantUser).Error
	return
}

func FindMerchantUserByIds(ids []uint) (merchantUsers []MerchantUser, err error) {
	err = db.Where("id IN ?", ids).Find(&merchantUsers).Error
	return
}

func FindMerchantUserByUserMerchant(userId, merchantId uint) (merchantUser MerchantUser, err error) {
	err = db.Where("merchant_id = ? AND user_id = ?", merchantId, userId).Find(&merchantUser).Error
	return
}

func FindAvailableMerchantUserWithMerchantByUser(userId uint) (merchantUsers []MerchantUser, err error) {
	err = db.Preload("Merchant").Where("user_id = ? AND status = ?", userId, "available").Find(&merchantUsers).Error
	return
}

func GetMerchantUserCountByMerchantStatus(merchantId uint, status string) uint {
	var count int64

	query := db.Table("merchant_users").Where("merchant_id = ?", merchantId)

	if status != "all" {
		query.Where("status = ?", status)
	}

	query.Count(&count)

	return uint(count)
}

func FindPaginationMerchantUserWithUserByStatus(merchantId uint, status string, offset int, size int) (merchantUsers []MerchantUser, err error) {
	query := db.Table("merchant_users").Where("merchant_id = ?", merchantId)

	if status != "all" {
		query.Where("status = ?", status)
	}

	err = query.Preload("User").Limit(size).Offset(offset).Order("id desc").Find(&merchantUsers).Error
	return
}

func UpdateMerchantUserRoleByIds(ids []uint, role string) (err error) {
	return db.Model(&MerchantUser{}).Where("id IN ?", ids).Update("role", role).Error
}

func DeleteMerchantUserById(id uint) (err error) {
	return db.Delete(&MerchantUser{}, id).Error
}

func SaveMerchantUser(merchantUser *MerchantUser) (err error) {
	return db.Save(merchantUser).Error
}
