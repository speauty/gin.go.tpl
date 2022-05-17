package bootstrap

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/inject"
	"github.com/gin-gonic/gin"
)

func init() {
	app := &App{}
	server := &Server{}

	container := &lib.Container{}
	err := container.Register(
		&inject.Object{Value: server},
		&inject.Object{Value: app},
	)
	if err != nil {
		panic(err)
	}
}

type App struct {
	Server Server `inject:"inline"`
}

func (app App) setGin() {
	gin.SetMode(gin.ReleaseMode)
}

func (app App) Run() {
	app.setGin()
	err := app.Server.NewServer().ListenAndServe()
	if err != nil {
		panic(err)
	}
}
