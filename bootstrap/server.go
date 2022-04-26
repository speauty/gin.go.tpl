package bootstrap

import (
	"gin.go.tpl/router"
	"net/http"
)

type Server struct{}

func (s Server) NewServer() *http.Server {
	return &http.Server{
		Addr:    ":80",
		Handler: router.Router{}.GetRouters(),
	}
}
