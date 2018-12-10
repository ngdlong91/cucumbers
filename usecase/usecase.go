package usecase

import (
	"github.com/astaxie/beego/context"
	"github.com/gin-gonic/gin"
	"github.com/ngdlong91/cucumbers/dto"
)

type UseCase interface {
	BeforeProcess()
	Process()
	AfterProcess()
	SetMoveNext(goNext bool)
	IsMoveNext() bool
}

type GinUseCase interface {
	SetContext(c *gin.Context)
}

type BeeUseCase interface {
	SetContext(ctx *context.Context)
	Response() dto.Response
	SetResponse(response dto.Response)
}

type InternalWorker func()

type Worker func(c *gin.Context)

func Run(uc UseCase) {
	uc.BeforeProcess()
	if !uc.IsMoveNext() {
		return
	}
	uc.Process()
	if !uc.IsMoveNext() {
		return
	}
	uc.AfterProcess()
}
