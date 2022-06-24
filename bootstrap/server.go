package bootstrap

import (
	"gin.go.tpl/kernel/cfg"
	"gin.go.tpl/kernel/log"
	"gin.go.tpl/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct{}

func (s Server) NewServer(engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    cfg.NewCfgApi("").Server.GetAddr(),
		Handler: router.Router{}.GetRouters(engine),
	}
}

func (s Server) StartServer(server *http.Server) {
	tmpCfg := cfg.NewCfgApi("")
	tmpLog := log.NewLogApi(nil)
	if tmpCfg.Server.IsHttp() {
		tmpLog.Info(tmpCfg.Server.Protocol, "server:", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	} else if tmpCfg.Server.IsHttps() {
		if tmpCfg.Server.SSL.Certificate == "" || tmpCfg.Server.SSL.CertificateKey == "" {
			panic("current server is https, no ssl configurated, you must to fix it quickly")
		}
		tmpLog.Info(tmpCfg.Server.Protocol, "server:", server.Addr)
		err := server.ListenAndServeTLS(tmpCfg.Server.SSL.Certificate, tmpCfg.Server.SSL.CertificateKey)
		if err != nil {
			panic(err)
		}
	} else {
		panic("no server configurated")
	}
	tmpLog.Info("server shutdown")
}
