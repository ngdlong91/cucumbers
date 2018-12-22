package ierr

import "github.com/ngdlong91/cucumbers/v2/dto"

type CustomError interface {
	IsSuccess() bool
	Msg() string
	Val() int
	Response(data interface{}) dto.Response
	SetMsg(msg string)
	Success() CustomError
}
