package db

const (
	VCStatusCreated    = "Created"
	VCStatusActive     = "Active"
	VCStatusProcessing = "Processing"
	VCStatusWithdraw   = "Withdraw"
	VCStatusInvalid    = "Invalid"
)

type VC struct {
	CommonModel

	PaymentId uint     `gorm:"index:idx_payment_id;not null" json:"payment_id"`
	Payment   *Payment `gorm:"foreignKey:PaymentId" json:"payment"`

	MerchantId uint      `gorm:"index:idx_merchant_id;not null" json:"merchant_id"`
	Merchant   *Merchant `gorm:"foreignKey:MerchantId" json:"merchant"`

	VCID      string `gorm:"index:idx_vcid;size:50;not null" json:"vc_id"`
	VCContent []byte `gorm:"type:TEXT;" json:"vc_content"`
	VCStatus  string `gorm:"index:idx_vc_status;size:100;not null" json:"vc_status"`
}

func FindPaginationMerchantVC(merchantId uint, status string, offset int, size int, option ...Options) (vcs []*VC, err error) {
	query := useOptions(option...).Where("merchant_id = ? and vc_status = ?", merchantId, status)

	err = query.Limit(size).Offset(offset).Order("id desc").Find(&vcs).Error
	return
}

func FindPaginationMerchantVCByVCId(merchantId uint, vcId []string, status string, offset int, size int, option ...Options) (vcs []*VC, err error) {
	query := useOptions(option...).Where("merchant_id = ? and vc_status = ? and vc_id in ?", merchantId, status, vcId)

	err = query.Limit(size).Offset(offset).Order("id desc").Find(&vcs).Error
	return
}

func FindVCsByVCIdAndStatus(vcIds []string, status string, option ...Options) (vcs []*VC, err error) {
	err = useOptions(option...).Where("vc_id in ? and vc_status = ?", vcIds, status).Find(&vcs).Error
	return
}

func FindVCsById(ids []uint, option ...Options) (vcs []*VC, err error) {
	err = useOptions(option...).Where("id in ?", ids).Find(&vcs).Error
	return
}

func FindVCsByVCId(vcIds []string, option ...Options) (vcs []*VC, err error) {
	err = useOptions(option...).Where("vc_id in ?", vcIds).Find(&vcs).Error
	return
}

func FindVCById(id uint, option ...Options) (vc *VC, err error) {
	err = useOptions(option...).Where("id = ?", id).Find(&vc).Error
	return
}

func UpdateVCStatus(merchantId uint, vcids []string, fromStatus string, toStatus string) error {
	return db.Model(&VC{}).
		Where("merchant_id = ? and vc_status = ? and vc_id in ?", merchantId, fromStatus, vcids).
		Update("vc_status", toStatus).
		Error
}

func MarkVCReceived(merchantId uint, vcids []string) error {
	return db.Model(&VC{}).
		Where("merchant_id = ? and vc_status = ? and vc_id in ?", merchantId, VCStatusCreated, vcids).
		Select("VCStatus", "VCContent").
		Updates(VC{VCStatus: VCStatusActive, VCContent: nil}).
		Error
}

func UpdateVC(vc *VC, fromStatus string, option ...Options) error {
	tx := useOptions(option...).Model(&VC{}).Where("id = ? and vc_status = ?", vc.ID, fromStatus).Updates(vc)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return ErrLockTryAgain
	} else {
		return nil
	}
}

func SaveVC(vc *VC) error {
	return db.Save(vc).Error
}
