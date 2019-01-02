package handler

import (
	"github.com/gin-gonic/gin"
)

type GinHandler interface {
	BindHandler(gr *gin.RouterGroup) GinHandler
	AddHandler(sub GinHandler) GinHandler
	Path() string
}
