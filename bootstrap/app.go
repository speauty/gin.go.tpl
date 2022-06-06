package bootstrap

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/constant"
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

	app.Context.Log = log.LogAPI

	// setMode by config from ini
	if app.Context.Config.Gin.Mode != "" {
		gin.SetMode(app.Context.Config.Gin.Mode)
	}
}

func (app *App) setMiddleware() {
	app.Engine.Use(middleware.RecoverMiddleware{}.Broken())
	// 非发布模式, 使用日志中间件
	if app.Context.Config.Gin.Mode != constant.GinModeRelease {
		app.Engine.Use(middleware.LogMiddleware{}.GoThrough())
	}
	app.Engine.Use(middleware.CorsMiddleware{}.SetHeaders())
}

func (app *App) Run() {
	// 初始化上下文
	lib.NewContextAPI().Init("./")

	app.Context = lib.NewContextAPI()
	app.setGin()
	app.setMiddleware()

	// 执行数据库迁移
	err := service.MigratorService{}.SyncTables(app.Context)
	if err != nil {
		panic(err)
	}

	app.Server.StartServer(app.Context, app.Server.NewServer(app.Context, app.Engine))
}
