package db

import (
	"gin.go.tpl/lib/db"
	"gorm.io/gorm"
)

type DB struct {
}

func (d DB) Get() *gorm.DB {
	return db.DBAPI.GetDB()
}
