package db

import "veric-backend/logic/db/types"

type PaymentTx struct {
	CommonModel

	PaymentId *uint    `gorm:"index:idx_payment_id" json:"payment_id"`
	Payment   *Payment `gorm:"foreignKey:PaymentId" json:"payment"`

	TxAddress    string        `gorm:"size:50;not null" json:"tx_address"`
	Amount       *types.BigInt `gorm:"not null;default:0" json:"amount"`
	AmountInCent uint64        `gorm:"not null;default:0" json:"amount_in_cent"`
	Hash         string        `gorm:"size:200;uniqueIndex:idx_hash;not null"`
	EventData    []byte        `gorm:"type:BLOB;not null"`
}

func QueryPaymentTxByHash(hash string, options ...Options) (tx *PaymentTx, err error) {
	err = useOptions(options...).Where("hash = ?", hash).Find(&tx).Error
	return
}

func SavePaymentTx(paymentTx *PaymentTx, options ...Options) (err error) {
	return useOptions(options...).Save(paymentTx).Error
}
