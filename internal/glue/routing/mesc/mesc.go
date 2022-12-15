package mesc

import (
	"dating/internal/glue/routing"
	"dating/internal/handler/middleware"
	"dating/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.RouterGroup, handler rest.Mesc, authMiddleware middleware.AuthMiddleware) {
	oauthRoutes := []routing.Router{
		{
			Method:      "GET",
			Path:        "/countries",
			Handler:     handler.GetCountries,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "POST",
			Path:        "/countries",
			Handler:     handler.CreateCountry,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "DELETE",
			Path:        "/countries/:id",
			Handler:     handler.DeleteCountry,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "GET",
			Path:        "/states",
			Handler:     handler.GetStates,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "POST",
			Path:        "/states",
			Handler:     handler.CreateState,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "DELETE",
			Path:        "/states/:id",
			Handler:     handler.DeleteState,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "GET",
			Path:        "/ethnicities",
			Handler:     handler.GetEthnicities,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "POST",
			Path:        "/ethnicities",
			Handler:     handler.CreateEthnicity,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Method:      "DELETE",
			Path:        "/ethnicities/:id",
			Handler:     handler.DeleteEthnicity,
			Middlewares: []gin.HandlerFunc{},
		},
	}
	routing.RegisterRoutes(router, oauthRoutes)
}
