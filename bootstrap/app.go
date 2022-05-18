package bootstrap

import (
	"gin.go.tpl/lib"
	libConfig "gin.go.tpl/lib/config"
	"github.com/gin-gonic/gin"
)

func init() {
}

type App struct {
	Server Server
	Config libConfig.Config
}

func (app App) setGin() {
	gin.SetMode(gin.ReleaseMode)
}

func (app App) Run() {
	app.setGin()
	// 上下文初始化
	lib.NewContextAPI().Init("./")

	err := app.Server.NewServer().ListenAndServe()
	if err != nil {
		panic(err)
	}
}
