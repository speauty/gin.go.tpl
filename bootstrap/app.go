package bootstrap

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/cache"
	"gin.go.tpl/lib/config"
	"gin.go.tpl/lib/constant"
	"gin.go.tpl/lib/db"
	"gin.go.tpl/lib/log"
	"gin.go.tpl/middleware"
	"gin.go.tpl/service"
	"github.com/gin-gonic/gin"
)

func init() {
}

type App struct {
	Engine  *gin.Engine
	Context *lib.Context
	Server  Server
}

func (app *App) setGin() {
	app.Engine = gin.Default()

	// setMode by config from ini
	if app.Context.Config.Gin.Mode != "" {
		gin.SetMode(app.Context.Config.Gin.Mode)
	}
}

func (app *App) setMiddleware() {
	app.Engine.Use(middleware.RecoverMiddleware{}.Exec())
	// 非发布模式, 使用日志中间件
	if app.Context.Config.Gin.Mode != constant.GinModeRelease {
		app.Engine.Use(middleware.LogMiddleware{}.Exec())
	}
	app.Engine.Use(middleware.CorsMiddleware{}.Exec(), middleware.LimiterMiddleware{}.Exec())
}

func (app *App) Run() {
	// 初始化 加载配置
	config.NewConfigApi("./")
	//@todo the MySql or PgSql config can't load at database node with viper, so using set
	config.ConfigApi.Database.MySql = config.ConfigApi.MySql

	// 初始化 日志&数据库&缓存
	log.NewLogApi(config.ConfigApi.Log)
	db.NewDbApi(config.ConfigApi.Database)
	cache.NewCacheApi(config.ConfigApi.Redis)

	// 初始化 上下文
	lib.NewContextApi().Init()

	app.Context = lib.NewContextApi()
	app.setGin()
	app.setMiddleware()

	// 执行数据库迁移
	err := service.MigratorService{}.SyncTables(app.Context)
	if err != nil {
		panic(err)
	}

	app.Server.StartServer(app.Context, app.Server.NewServer(app.Context, app.Engine))
}
