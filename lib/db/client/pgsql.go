package client

import "gorm.io/gorm"

// if you want to use pgsql, should exec gorm.io/driver/postgres

type PgSqlClient struct{}

func (mc PgSqlClient) OpenDSN(dsn string) gorm.Dialector {
	//@todo wait to complete
	// return postgres.Open(dsn)
	return nil
}
