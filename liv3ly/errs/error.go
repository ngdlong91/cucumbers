package ierr

import "github.com/ngdlong91/cucumbers/liv3ly/dto"

type CustomError interface {
	IsSuccess() bool
	Msg() string
	Val() int
	Response(data interface{}) dto.Response
	SetMsg(msg string)
	Success() CustomError
	// Tracking
	//Debug() bool
}
