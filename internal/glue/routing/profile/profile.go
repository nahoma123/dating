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

		// update profile
		{
			Method:      "PATCH",
			Path:        "/profiles/:id",
			Handler:     handler.UpdateProfile,
			Middlewares: []gin.HandlerFunc{},
		},
		// get profile
		{
			Method:      "GET",
			Path:        "/profiles/:id",
			Handler:     handler.GetProfile,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "GET",
			Path:        "/customers",
			Handler:     handler.GetCustomers,
			Middlewares: []gin.HandlerFunc{authMiddleware.BindUser("15502629-f754-42e1-8475-3a3915a4be52")},
		},
		{
			Method:      "GET",
			Path:        "/customers/discover",
			Handler:     handler.DiscoverNewUsers,
			Middlewares: []gin.HandlerFunc{authMiddleware.BindUser("15502629-f754-42e1-8475-3a3915a4be52")},
		},
	}
	routing.RegisterRoutes(router, profileRoutes)
}
