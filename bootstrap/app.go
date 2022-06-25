package bootstrap

import (
	"gin.go.tpl/kernel/cache"
	"gin.go.tpl/kernel/cfg"
	"gin.go.tpl/kernel/constant"
	"gin.go.tpl/kernel/db"
	"gin.go.tpl/kernel/log"
	"gin.go.tpl/middleware"
	"gin.go.tpl/service"
	"github.com/gin-gonic/gin"
)

func init() {
}

type App struct {
	Engine *gin.Engine
	Server Server
}

func (app *App) setGin() {
	app.Engine = gin.Default()

	// setMode by config from ini
	if cfg.NewCfgApi("").Gin.Mode != "" {
		gin.SetMode(cfg.NewCfgApi("").Gin.Mode)
	}
}

func (app *App) setMiddleware() {
	app.Engine.Use(middleware.RecoverMiddleware{}.Exec())
	// 非发布模式, 使用日志中间件
	if cfg.NewCfgApi("").Gin.Mode != constant.GinModeRelease {
		app.Engine.Use(middleware.LogMiddleware{}.Exec())
	}
	app.Engine.Use(
		middleware.CorsMiddleware{}.Exec(), middleware.LimiterMiddleware{}.Exec(),
	)
}

func (app *App) Run() {
	// 初始化 加载配置
	tmpCfg := cfg.NewCfgApi("./")
	//@todo the MySql or PgSql config can't load at database node with viper, so using set
	tmpCfg.Database.MySql = tmpCfg.MySql

	// 初始化 日志&数据库&缓存
	log.NewLogApi(&tmpCfg.Log)
	db.NewDbApi(&tmpCfg.Database)
	cache.NewCacheApi(&tmpCfg.Redis)

	app.setGin()
	app.setMiddleware()

	// 执行数据库迁移
	err := service.MigratorService{}.SyncTables()
	if err != nil {
		panic(err)
	}

	app.Server.StartServer(app.Server.NewServer(app.Engine))
}
