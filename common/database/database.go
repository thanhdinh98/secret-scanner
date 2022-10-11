package database

import (
	"sync"
	"time"

	"guardian/common"
	"guardian/common/rootcloser"
	"guardian/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	lock = &sync.Mutex{}
)

func Initialize() {
	if extension == nil {
		defer lock.Unlock()
		lock.Lock()
		if extension == nil {
			db, err := gorm.Open(postgres.Open(config.KeyDBConnection), &gorm.Config{
				SkipDefaultTransaction: true,
			})
			common.PanicOnError(err)

			sqlDB, err := db.DB()
			common.PanicOnError(err)

			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(time.Hour)

			extension = &Extension{
				db,
			}

			rootcloser.RegisterCloser(closeDB(extension))
		}
	}
}

func closeDB(ext *Extension) func() {
	return func() {
		if nativeSqlDB, err := ext.db.DB(); err == nil {
			nativeSqlDB.Close()
		}
	}
}

func GetDB() LightHandler {
	return getExtension()
}
