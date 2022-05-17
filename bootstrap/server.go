package bootstrap

import (
	"gin.go.tpl/router"
	"net/http"
)

type Server struct{}

func (s Server) NewServer() *http.Server {
	return &http.Server{
		Addr:    "127.0.0.1:10086",
		Handler: router.Router{}.GetRouters(),
	}
}
