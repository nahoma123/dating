package profile

import (
	"dating/internal/glue/routing"
	"dating/internal/handler/middleware"
	"dating/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.RouterGroup, handler rest.Profile, authMiddleware middleware.AuthMiddleware) {
	profileRoutes := []routing.Router{
		{
			Method:      "POST",
			Path:        "/profiles/register",
			Handler:     handler.Register,
			Middlewares: []gin.HandlerFunc{},
		},
	}
	routing.RegisterRoutes(router, profileRoutes)
}
