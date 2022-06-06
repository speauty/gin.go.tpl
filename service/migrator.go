package service

import (
	"gin.go.tpl/db"
	"gin.go.tpl/db/entity"
	"gin.go.tpl/lib"
)

type MigratorService struct{}

func (ms MigratorService) SyncTables(ctx *lib.Context) error {
	tmpDb := db.DB{}
	if tmpDb.GetConfigIsMigration() {
		err := tmpDb.Get().Set("gorm:table_options", "ENGINE=InnoDB").
			AutoMigrate(
				&entity.User{},
			)
		if err != nil {
			ctx.Log.Error(err)
			return err
		}
	}
	return nil
}
