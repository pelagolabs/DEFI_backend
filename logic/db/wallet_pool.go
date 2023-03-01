package db

type WalletPool struct {
	CommonModel

	ChainName string `gorm:"index:idx_search;size:50;not null" json:"chain_name"`
	Addr      string `gorm:"index:idx_search;size:100;not null" json:"addr"`
}

func AllWalletFromPool() (wallets []*WalletPool, err error) {
	err = db.
		Find(&wallets).
		Error
	return
}

func QueryWalletUseChainName(chainName string) (wallets []*WalletPool, err error) {
	err = db.
		Where("chain_name = ?", chainName).
		Find(&wallets).
		Error
	return
}
