package bootstrap

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/middleware"
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
	app.Engine.Use(middleware.CorsMiddleware{}.SetHeaders())
}

func (app *App) Run() {
	// to initialize context
	lib.NewContextAPI().Init("./")

	app.Context = lib.NewContextAPI()
	app.setGin()
	app.setMiddleware()

	app.Server.StartServer(app.Context, app.Server.NewServer(app.Context, app.Engine))
}
