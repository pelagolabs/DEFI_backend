package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"time"
	"veric-backend/internal/log"
	"veric-backend/logic/config"
)

var db *gorm.DB

type CommonModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func buildDSN() string {
	dbConfig := config.Get().DB
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
		dbConfig.TimeZone,
	)
}

func mustSuccess(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	log.GetLogger().Info("init mysql db...")

	mainDB, err := gorm.Open(mysql.Open(buildDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = mainDB
	mainDB.Logger = zapgorm2.New(log.GetLogger())

	if config.Get().Debug.Verbose {
		db = db.Debug()
	}

	migrate := db

	mustSuccess(migrate.AutoMigrate(&Tx{}))
	mustSuccess(migrate.AutoMigrate(&VC{}))
	mustSuccess(migrate.AutoMigrate(&ProcessItem{}))
	mustSuccess(migrate.AutoMigrate(&Application{}))
	mustSuccess(migrate.AutoMigrate(&Blockchain{}))
	mustSuccess(migrate.AutoMigrate(&Currency{}))
	mustSuccess(migrate.AutoMigrate(&Manager{}))
	mustSuccess(migrate.AutoMigrate(&Merchant{}))
	mustSuccess(migrate.AutoMigrate(&Payment{}))
	mustSuccess(migrate.AutoMigrate(&PaymentTx{}))
	mustSuccess(migrate.AutoMigrate(&Permission{}))
	mustSuccess(migrate.AutoMigrate(&Pool{}))
	mustSuccess(migrate.AutoMigrate(&MerchantUser{}))
	mustSuccess(migrate.AutoMigrate(&User{}))
	mustSuccess(migrate.AutoMigrate(&Withdraw{}))
	mustSuccess(migrate.AutoMigrate(&Task{}))
	mustSuccess(migrate.AutoMigrate(&FeeWithdrawLog{}))
	mustSuccess(migrate.AutoMigrate(&PaymentFee{}))
	mustSuccess(migrate.AutoMigrate(&WalletPool{}))
	mustSuccess(migrate.AutoMigrate(&PaymentBalance{}))
}
