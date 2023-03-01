package db

type Currency struct {
	CommonModel

	Name string `gorm:"size:100;not null" json:"name"`

	Symbol               string     `gorm:"size:50;not null" json:"symbol"`
	SearchName           string     `gorm:"size:100;not null" json:"search_name"`
	EventBlock           uint64     `gorm:"not null" json:"event_block"`
	ChainId              uint       `gorm:"not null;default:1" json:"chain_id"`
	Blockchain           Blockchain `gorm:"foreignKey:ChainId" json:"chain_detail"`
	DecimalCount         uint       `gorm:"not null;default:18" json:"decimal_count"`
	FriendlyDecimalCount uint       `gorm:"not null;default:16" json:"friendly_decimal_count"`
	ContractAddress      string     `gorm:"size:100;not null;default:''" json:"contract_address"`
	ImageUrl             string     `gorm:"size:200;not null;default:''" json:"image_url"`
	IsNative             bool       `gorm:"type:tinyint;not null;default:0" json:"is_native"`
	Status               string     `gorm:"size:50;not null" json:"status"`
	IsHidden             bool       `gorm:"type:tinyint;not null;default:0" json:"is_hidden"`
}

func FindAllCurrencyByChain(chainId uint, options ...Options) (currencies []*Currency, err error) {
	err = useOptions(options...).Where("chain_id = ?", chainId).Find(&currencies).Error
	return
}

func FindCurrencyByChainAndSymbol(chainId uint, symbol string, options ...Options) (currency *Currency, err error) {
	err = useOptions(options...).Where("chain_id = ? and symbol = ?", chainId, symbol).Find(&currency).Error
	return
}

func FindAvailableCurrencies(options ...Options) (currencies []Currency, err error) {
	err = useOptions(options...).Where("status = ?", "available").Find(&currencies).Error
	return
}

func FindAvailableNotHiddenCurrencies(options ...Options) (currencies []Currency, err error) {
	err = useOptions(options...).Where("status = ? AND is_hidden = ?", "available", false).Find(&currencies).Error
	return
}

func FindAvailableCurrenciesById(ids []uint) (currencies []Currency, err error) {
	err = db.Where("id IN ? AND status = ?", ids, "available").Find(&currencies).Error
	return
}

func FindAvailableCurrencyById(id uint, options ...Options) (currency *Currency, err error) {
	err = useOptions(options...).Where("id = ? AND status = ?", id, "available").Find(&currency).Error
	return
}

func FindCurrencyById(id uint, options ...Options) (currency *Currency, err error) {
	err = useOptions(options...).Where("id = ?", id).Find(&currency).Error
	return
}

func FindAllNativeCurrency(options ...Options) (currencies []*Currency, err error) {
	err = useOptions(options...).Where("is_native = 1").Find(&currencies).Error
	return
}

func GetCurrencyCountByStatus(status string) uint {
	var count int64

	query := db.Table("currencies")

	if status != "all" {
		query.Where("status = ?", status)
	}

	query.Where("is_hidden = ?", false).Count(&count)

	return uint(count)
}

func FindPaginationCurrencyByStatus(status string, offset int, size int) (currencies []Currency, err error) {
	query := db.Table("currencies")

	if status != "all" {
		query.Where("status = ?", status)
	}

	err = query.Where("is_hidden = ?", false).Limit(size).Offset(offset).Order("id desc").Find(&currencies).Error
	return
}

func UpdateCurrencyEventBlock(id uint, eventBlock uint64) (err error) {
	return db.Model(&Currency{}).Where("id = ?", id).Update("event_block", eventBlock).Error
}

func SaveCurrency(currency *Currency) (err error) {
	return db.Save(currency).Error
}
