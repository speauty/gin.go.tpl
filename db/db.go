package db

import (
	"gin.go.tpl/lib/db"
	"gorm.io/gorm"
)

type DB struct {
}

func (d DB) GetApi() *db.DB {
	return db.DBApi
}

func (d DB) Get() *gorm.DB {
	return d.GetApi().GetDB()
}

func (d DB) GetConfigIsMigration() bool {
	return d.GetApi().GetConfig().IsMigration
}

func (d DB) Create(value interface{}) error {
	return d.Get().Create(value).Error
}
