package rest

import "github.com/gin-gonic/gin"

type OAuth interface {
	Test(ctx *gin.Context)
}
