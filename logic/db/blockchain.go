package db

import "errors"

type Blockchain struct {
	CommonModel

	Name         string `gorm:"size:100;not null" json:"name"`
	Endpoint     string `gorm:"size:150;not null" json:"endpoint"`
	ContractName string `gorm:"size:30;not null;default:''" json:"contract_name"`
	ImageUrl     string `gorm:"size:200;not null;default:''" json:"image_url"`
	Status       string `gorm:"index:idx_blockchain_status;size:100;not null" json:"status"`
	IsHidden     bool   `gorm:"type:tinyint;not null;default:0" json:"is_hidden"`
	FarmAddress  string `gorm:"size:100;not null;default:''" json:"farm_address"`
}

func AllBlockchain() (blockchains []*Blockchain, err error) {
	err = db.Find(&blockchains).Error
	return
}

func FindBlockchainById(id uint) (blockchain *Blockchain, err error) {
	err = db.Where("id = ?", id).Find(&blockchain).Error
	return
}

func FindBlockchainByContractName(name string) (blockchain *Blockchain, err error) {
	err = db.Where("contract_name = ?", name).Find(&blockchain).Error
	return
}

func GetBlockchainCountByStatus(status string) uint {
	var count int64

	query := db.Table("blockchains")

	if status != "all" {
		query.Where("status = ?", status)
	}

	query.Where("is_hidden = ?", false).Count(&count)

	return uint(count)
}

func FindPaginationBlockchainByStatus(status string, offset int, size int) (blockchains []Blockchain, err error) {
	query := db.Table("blockchains")

	if status != "all" {
		query.Where("status = ?", status)
	}

	err = query.Where("is_hidden = ?", false).Limit(size).Offset(offset).Order("id desc").Find(&blockchains).Error
	return
}

func SaveBlockchain(blockchain *Blockchain) (err error) {
	return db.Save(blockchain).Error
}

func SaveBlockchainWithStatusLock(blockchain *Blockchain, oriStatus string, options ...Options) (err error) {
	tx := useOptions(options...).Model(&Blockchain{}).Where("id = ? and status = ?", blockchain.ID, oriStatus).Updates(blockchain)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return errors.New("ori status check fail")
	} else {
		return nil
	}
}
