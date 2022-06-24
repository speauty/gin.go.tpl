package service

import (
	"gin.go.tpl/db"
	"gin.go.tpl/db/entity"
	"gin.go.tpl/kernel/log"
)

type MigratorService struct{}

func (ms MigratorService) SyncTables() error {
	tmpDb := db.DB{}
	if tmpDb.GetConfigIsMigration() {
		err := tmpDb.Get().Set("gorm:table_options", "ENGINE=InnoDB").
			AutoMigrate(
				&entity.User{},
			)
		if err != nil {
			log.NewLogApi(nil).Error(err)
			return err
		}
	}
	return nil
}
