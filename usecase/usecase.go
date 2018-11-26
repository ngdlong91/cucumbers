package usecase

import (
	"github.com/gin-gonic/gin"
)

type UseCase interface {
	BeforeProcess()
	Process()
	AfterProcess()
	SetMoveNext(goNext bool)
	IsMoveNext() bool
	SetContext(c *gin.Context)
}

type InternalUseCase interface {
	BeforeProcess()
	Process()
	AfterProcess()
	SetMoveNext(goNext bool)
	IsMoveNext() bool
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
