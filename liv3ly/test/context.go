package test

import "github.com/gin-gonic/gin"

/**
Test context struct for easy test case declaration
Ex:
	ctx := test.NewGinContext().Set("test", 5)
	value := ctx.GetGinContext().GetInt("test")
	fmt.Println("Value ", value)
 */
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

func (c *Context) GinSet(key string, data interface{}) *Context {
	c.ginContext.Set(key, data)
	return c
}


