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
	SetContext(ctx *gin.Context)
}

type BeeUseCase interface {
	SetContext(ctx *context.Context)
	Response() dto.Response
	SetResponse(resp dto.Response)
}

type InternalWorker func()

type Worker func(c *gin.Context)

func Run(uc UseCase) {
	uc.SetMoveNext(false)
	uc.BeforeProcess()
	if !uc.IsMoveNext() {
		return
	}
	uc.SetMoveNext(false)
	uc.Process()
	if !uc.IsMoveNext() {
		return
	}
	uc.AfterProcess()
}
