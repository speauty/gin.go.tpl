package db

import (
	"gin.go.tpl/kernel/cfg"
	"gin.go.tpl/kernel/constant"
	"gin.go.tpl/kernel/db/client"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
	"time"
)

var (
	api  *Db
	once sync.Once
)

func NewDbApi(cfg *cfg.DatabaseConf) *Db {
	once.Do(func() {
		api = &Db{cfg: cfg}
		api.init()
	})
	return api
}

type IDb interface {
	GetDriver() string
	GetCfg() *cfg.DatabaseConf
	GetDb() *gorm.DB
	init()
}

type Db struct {
	cfg           *cfg.DatabaseConf
	currentDB     *gorm.DB
	currentDriver string
}

func (d *Db) GetDriver() string {
	return d.currentDriver
}

func (d *Db) GetCfg() *cfg.DatabaseConf {
	return d.cfg
}

func (d *Db) GetDb() *gorm.DB {
	if d.currentDB == nil {
		d.init()
	}
	return d.currentDB
}

func (d *Db) init() {
	d.setConn()
	d.setPool()
}

func (d *Db) setConn() {
	if d.currentDB == nil {
		var err error
		if d.cfg.DBDriver == constant.DatabaseDriverMysql {
			d.currentDB, err = gorm.Open(client.MySqlClient{}.OpenDSN(d.cfg.MySql.DSN), d.getOption())
			d.currentDriver = constant.DatabaseDriverMysql
		} else if d.cfg.DBDriver == constant.DatabaseDriverPgsql {
			d.currentDB, err = gorm.Open(client.PgSqlClient{}.OpenDSN(d.cfg.PgSql.DSN), d.getOption())
			d.currentDriver = constant.DatabaseDriverPgsql
		} else {
			panic("the database driver not set, or not supported")
		}
		if err != nil {
			panic(err)
		}
	}
}

func (d *Db) setPool() {
	if d.cfg.IsPool > 0 {
		sqlDB, _ := d.GetDb().DB()
		sqlDB.SetMaxIdleConns(d.cfg.MaxIdleConn)
		sqlDB.SetMaxOpenConns(d.cfg.MaxOpenConn)
		sqlDB.SetConnMaxLifetime(time.Duration(d.cfg.MaxLifetime))
	}
}

func (d *Db) getOption() *gorm.Config {
	gormConfig := &gorm.Config{}
	// to disable default transaction
	gormConfig.SkipDefaultTransaction = d.cfg.SkipDefaultTransaction
	// to set a logger for gorm
	gormConfig.Logger = logger.Default.LogMode(logger.Warn)

	gormConfig.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   d.cfg.PrefixTable,
		SingularTable: d.cfg.SingularTable,
		NoLowerCase:   d.cfg.NoLowerCase,
	}

	return gormConfig
}
