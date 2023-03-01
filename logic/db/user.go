package db

import "errors"

const (
	DidUpStatusCreated = "created"
	DidUpStatusSuccess = "success"
	DidUpStatusPending = "pending"
	DidUpStatusFailed  = "failed"
)

type User struct {
	CommonModel

	Address     string `gorm:"size:50;not null;uniqueIndex:idx_user_address" json:"address"`
	Did         string `gorm:"size:100;not null;uniqueIndex:idx_user_did" json:"did"`
	Nickname    string `gorm:"size:100;not null" json:"nickname"`
	Status      string `gorm:"index:idx_user_status;size:50;not null" json:"status"`
	DidPubKey   string `gorm:"size:500;not null;default:''" json:"did_pub_key"`
	DidUpStatus string `gorm:"size:30;not null;index:idx_user_did_status" json:"did_up_status"`
	DidUpTx     string `gorm:"size:100;not null;default:''" json:"did_up_tx"`
}

func FindUserByAddress(address string) (user *User, err error) {
	err = db.Where("address = ?", address).Find(&user).Error
	return
}

func FindUserByDid(did string) (user *User, err error) {
	err = db.Where("did = ?", did).Find(&user).Error
	return
}

func GetUserByDidUpStatus(didUpStatus string) (users []*User, err error) {
	err = db.Where("did_up_status = ?", didUpStatus).Find(&users).Error
	return
}

func SaveUser(user *User) (err error) {
	return db.Save(user).Error
}

func SaveUserWithDidStatusLock(user *User, oriDidStatus string, options ...Options) (err error) {
	tx := useOptions(options...).Model(&User{}).Where("id = ? and did_up_status = ?", user.ID, oriDidStatus).Updates(user)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return errors.New("ori status check fail")
	} else {
		return nil
	}
}
