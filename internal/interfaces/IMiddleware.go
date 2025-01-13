package interfaces

import "github.com/gin-gonic/gin"

type IMiddlewareHandler interface {
	MiddlewareValidateAuth(c *gin.Context)
	MiddlewareRefreshToken(ctx *gin.Context)
}
