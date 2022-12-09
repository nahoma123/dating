package auth

import (
	"dating/internal/glue/routing"
	"dating/internal/handler/middleware"
	"dating/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.RouterGroup, handler rest.OAuth, authMiddleware middleware.AuthMiddleware) {
	oauthRoutes := []routing.Router{
		{
			Method:      "GET",
			Path:        "/test",
			Handler:     handler.Test,
			Middlewares: []gin.HandlerFunc{},
		},
	}
	routing.RegisterRoutes(router, oauthRoutes)
}
