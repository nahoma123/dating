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
			Middlewares: []gin.HandlerFunc{authMiddleware.BindUser("cc66a19c-2f8b-400d-9051-5122573a9974")},
		},
	}
	routing.RegisterRoutes(router, profileRoutes)
}
