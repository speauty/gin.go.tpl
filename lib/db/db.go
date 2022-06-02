package db

import (
	"gin.go.tpl/lib/config"
	"gin.go.tpl/lib/constant"
	"gin.go.tpl/lib/db/client"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var (
	DBAPI  *DB
	DBOnce sync.Once
)

type DB struct {
	Config        config.DatabaseConf
	currentDB     *gorm.DB
	currentDriver string
}

func NewDBAPI(config config.DatabaseConf) *DB {
	DBOnce.Do(func() {
		DBAPI = &DB{Config: config}
		DBAPI.Init()
	})
	return DBAPI
}

func (d *DB) Init() {
	d.setConn()
	d.setPool()
}

func (d *DB) GetDB() *gorm.DB {
	if d.currentDB == nil {
		d.Init()
	}
	return d.currentDB
}

func (d DB) GetDriver() string {
	return d.currentDriver
}

func (d *DB) setConn() {
	if d.currentDB == nil {
		var err error
		if d.Config.DBDriver == constant.DatabaseDriverMysql {
			d.currentDB, err = gorm.Open(client.MySqlClient{}.OpenDSN(d.Config.MySql.DSN), d.getOption())
			d.currentDriver = constant.DatabaseDriverMysql
		} else if d.Config.DBDriver == constant.DatabaseDriverPgsql {
			d.currentDB, err = gorm.Open(client.PgSqlClient{}.OpenDSN(d.Config.PgSql.DSN), d.getOption())
			d.currentDriver = constant.DatabaseDriverPgsql
		} else {
			panic("the database driver not set, or not supported")
		}
		if err != nil {
			panic(err)
		}
	}
}

func (d DB) setPool() {
	if d.Config.IsPool > 0 {
		sqlDB, _ := d.GetDB().DB()
		sqlDB.SetMaxIdleConns(d.Config.MaxIdleConn)
		sqlDB.SetMaxOpenConns(d.Config.MaxOpenConn)
		sqlDB.SetConnMaxLifetime(time.Duration(d.Config.MaxLifetime))
	}
}

func (d DB) getOption() *gorm.Config {
	gormConfig := &gorm.Config{}
	// to disable default transaction
	gormConfig.SkipDefaultTransaction = d.Config.SkipDefaultTransaction
	// to set a logger for gorm
	gormConfig.Logger = logger.Default.LogMode(logger.Warn)

	gormConfig.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   d.Config.PrefixTable,
		SingularTable: d.Config.SingularTable,
		NoLowerCase:   d.Config.NoLowerCase,
	}

	return gormConfig
}
