package test

import "github.com/gin-gonic/gin"

type Context struct {
	ginContext *gin.Context
}

func NewGinContext() *Context {
	return &Context{
		ginContext: &gin.Context{},
	}
}

func (c *Context) GetGinContext() *gin.Context {
	return c.ginContext
}

func (c *Context) Set(key string, data interface{}) *Context {
	c.ginContext.Set(key, data)
	return c
}


