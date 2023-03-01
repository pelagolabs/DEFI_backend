package db

import "errors"

type Pool struct {
	CommonModel

	Name       string   `gorm:"size:100;not null" json:"name"`
	Describe   string   `gorm:"size:200" json:"describe"`
	CurrencyId uint     `gorm:"index:idx_pool_currency_id;not null" json:"-"`
	Currency   Currency `gorm:"foreignKey:CurrencyId" json:"currency"`
	Address    string   `gorm:"size:50;not null" json:"address"`
	Tvl        float64  `gorm:"not null;default:0" json:"tvl"`
	Apy        float64  `gorm:"not null;default:0" json:"apy"`
	Status     string   `gorm:"index:idx_pool_status;size:100;not null" json:"status"`
}

func FindPoolById(id uint) (pool *Pool, err error) {
	err = db.Where("id = ?", id).Find(&pool).Error
	return
}

func FindPoolWithCurrencyById(id uint) (pool *Pool, err error) {
	err = db.Preload("Currency").Where("id = ?", id).Find(&pool).Error
	return
}

func FindPoolWithCurrencyByStatus(status []string) (pools []*Pool, err error) {
	err = db.Preload("Currency.Blockchain").Where("status IN ?", status).Find(&pools).Error
	return
}

func FindPoolByCurrencyId(currencyId uint) (pool *Pool, err error) {
	err = db.Where("currency_id = ?", currencyId).Find(&pool).Error
	return
}

func GetPoolCountByStatus(status string) uint {
	var count int64

	query := db.Table("pools")

	if status != "all" {
		query.Where("status = ?", status)
	}

	query.Count(&count)

	return uint(count)
}

func FindPaginationPoolByStatus(status string, offset int, size int) (pools []Pool, err error) {
	query := db.Preload("Currency").Table("pools")

	if status != "all" {
		query.Where("status = ?", status)
	}

	err = query.Limit(size).Offset(offset).Order("id desc").Find(&pools).Error
	return
}

func SavePool(pool *Pool) (err error) {
	return db.Save(pool).Error
}

func SavePoolWithStatusLock(pool *Pool, oriStatus string, options ...Options) (err error) {
	tx := useOptions(options...).Model(&Pool{}).Where("id = ? and status = ?", pool.ID, oriStatus).Updates(pool)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return errors.New("ori status check fail")
	} else {
		return nil
	}
}
