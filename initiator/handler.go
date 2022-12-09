package initiator

import (
	"dating/internal/handler/rest"
	"dating/internal/handler/rest/oauth"
	"dating/platform/logger"
)

type Handler struct {
	// TODO implement
	oauth rest.OAuth
}

func InitHandler(module Module, log logger.Logger) Handler {
	return Handler{
		// TODO implement
		oauth: oauth.InitAuth(log, module.AuthModule),
	}
}
