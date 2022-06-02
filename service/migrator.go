package service

import (
	"gin.go.tpl/db"
	"gin.go.tpl/db/entity"
	"gin.go.tpl/lib"
)

type MigratorService struct{}

func (ms MigratorService) SyncTables(ctx *lib.Context) {
	err := db.DB{}.Get().Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&entity.User{},
		)
	if err != nil {
		ctx.Log.Error(err)
	}
}
