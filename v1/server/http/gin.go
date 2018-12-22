package http

import (
	"ngdlong91/klit/server/config"
	"ngdlong91/klit/server/handler"
	"ngdlong91/klit/storage"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GinServer struct {
	config  *config.GinConfig
	logger  *logrus.Entry
	storage storage.Storage
	cache   storage.Cache
	engine  *gin.Engine
}

func NewGinServer(config *config.GinConfig) *GinServer {
	return &GinServer{
		config: config,
		engine: gin.New(),
	}
}

func (s *GinServer) AddHandler(h handler.GinHandler) *GinServer {
	h.BindHandler(s.engine.Group(h.Path()))
	return s
}

func (s *GinServer) Run() {
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	s.engine.Run(addr)
}
