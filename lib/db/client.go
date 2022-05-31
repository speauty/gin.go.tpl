package db

import "gorm.io/gorm"

type Client interface {
	OpenDSN(dsn string) gorm.Dialector
}
