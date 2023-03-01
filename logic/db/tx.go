package db

import "veric-backend/logic/db/types"

const (
	TxStatusUnknown = "Unknown"
	TxStatusSuccess = "Success"
	TxStatusFail    = "Fail"
)

type Tx struct {
	CommonModel

	Name        string `gorm:"size:200;index:idx_name;not null"`
	Method      string `gorm:"size:100;index:idx_tx_method;not null"`
	TaskId      string `gorm:"size:64;not null"`
	Hash        string `gorm:"size:200;index:idx_hash;not null"`
	Status      string `gorm:"size:20;index:idx_name;not null"`
	ChainName   string `gorm:"size:20;not null"`
	GasUsed     *types.BigInt
	Why         *string `gorm:"type:TEXT;"`
	TxData      []byte  `gorm:"type:BLOB;not null"`
	ReceiptData []byte  `gorm:"type:TEXT;"`
}

func FindTxByHashAndTaskId(hash, taskId string) (tx *Tx, err error) {
	err = db.Where("hash = ? and task_id = ?", hash, taskId).Find(&tx).Error
	return
}

func FindTxByNameAndStatus(name, status string) (txs []*Tx, err error) {
	err = db.Where("name = ? and status = ?", name, status).Find(&txs).Error
	return
}

func FindTxByCreateTime(start, end string) (txs []*Tx, err error) {
	query := db.Table("txes")

	if start != "all" {
		query.Where("created_at BETWEEN ? AND ?", start, end)
	}

	err = query.Find(&txs).Error

	return
}

func GetTxCountByStatusMethodDate(status, method string, date string) uint {
	var count int64

	query := db.Table("txes")

	if status != "all" {
		query.Where("status = ?", status)
	}

	if method != "all" {
		query.Where("method = ?", method)
	}

	if date != "" {
		query.Where("created_at BETWEEN ? AND ?", date+" 00:00:00", date+" 23:59:59")
	}

	query.Count(&count)

	return uint(count)
}

func FindPaginationTxByStatusMethodDate(status, method, date string, offset int, size int) (txs []*Tx, err error) {
	query := db.Table("txes")

	if status != "all" {
		query.Where("status = ?", status)
	}

	if method != "all" {
		query.Where("method = ?", method)
	}

	if date != "" {
		query.Where("created_at BETWEEN ? AND ?", date+" 00:00:00", date+" 23:59:59")
	}

	err = query.Limit(size).Offset(offset).Order("id desc").Find(&txs).Error
	return
}

func SaveTx(tx *Tx) (err error) {
	return db.Save(tx).Error
}
