package client

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlClient struct{}

func (mc MySqlClient) OpenDSN(dsn string) gorm.Dialector {
	return mysql.Open(dsn)
}
