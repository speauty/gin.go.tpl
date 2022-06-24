package db

import (
	"gin.go.tpl/kernel/db"
	"gorm.io/gorm"
)

type DB struct {
}

func (d DB) GetApi() *db.Db {
	return db.NewDbApi(nil)
}

func (d DB) Get() *gorm.DB {
	return d.GetApi().GetDb()
}

func (d DB) GetConfigIsMigration() bool {
	return d.GetApi().GetCfg().IsMigration
}

func (d DB) Create(value interface{}) error {
	return d.Get().Create(value).Error
}
