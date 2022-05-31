package bootstrap

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct{}

func (s Server) NewServer(ctx *lib.Context, engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    lib.NewContextAPI().Config.Server.GetAddr(),
		Handler: router.Router{}.GetRouters(ctx, engine),
	}
}

func (s Server) StartServer(ctx *lib.Context, server *http.Server) {
	if ctx.Config.Server.IsHttp() {
		ctx.Log.Info(ctx.Config.Server.Protocol, "server:", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	} else if ctx.Config.Server.IsHttps() {
		if ctx.Config.Server.SSL.Certificate == "" || ctx.Config.Server.SSL.CertificateKey == "" {
			panic("current server is https, no ssl configurated, you must to fix it quickly")
		}
		ctx.Log.Info(ctx.Config.Server.Protocol, "server:", server.Addr)
		err := server.ListenAndServeTLS(ctx.Config.Server.SSL.Certificate, ctx.Config.Server.SSL.CertificateKey)
		if err != nil {
			panic(err)
		}
	} else {
		panic("no server configurated")
	}
	ctx.Log.Info("server shutdown")
}
