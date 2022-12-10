package initiator

import (
	// swager docs import
	// "dating/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"

	"dating/internal/glue/routing/auth"
	"dating/internal/glue/routing/profile"
	"dating/internal/handler/middleware"
	"dating/platform/logger"
)

func InitRouter(router *gin.Engine, group *gin.RouterGroup, handler Handler, module Module, log logger.Logger, publicKeyPath string) {

	authMiddleware := middleware.InitAuthMiddleware(module.AuthModule, nil)

	docs.SwaggerInfo.BasePath = "/v1"

	// swager docs import
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth.InitRoute(group, handler.oauth, authMiddleware)

	profile.InitRoute(group, handler.profile, authMiddleware)

}
