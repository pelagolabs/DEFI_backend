package db

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ErrLockTryAgain = errors.New("optimistic locking failed, try again...")
)

type Options func(*gorm.DB) *gorm.DB

func WithSqlDebug(ori *gorm.DB) *gorm.DB {
	return ori.Debug()
}

func WithTx(f func(tx Options) error, options ...Options) error {
	tx := useOptions(options...).Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	err := f(func(_ *gorm.DB) *gorm.DB {
		return tx
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func SavePoint(pointName string, options ...Options) error {
	return useOptions(options...).SavePoint(pointName).Error
}

func RollbackTo(pointName string, options ...Options) error {
	return useOptions(options...).RollbackTo(pointName).Error
}

func WithSelect(fields []string) Options {
	return func(g *gorm.DB) *gorm.DB {
		return g.Select(fields)
	}
}

func WithPreload(name string) Options {
	return func(g *gorm.DB) *gorm.DB {
		return g.Preload(name)
	}
}

func useOptions(options ...Options) *gorm.DB {
	var optionsDB = db
	for _, option := range options {
		optionsDB = option(optionsDB)
	}
	return optionsDB
}
