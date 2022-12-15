package initiator

import (
	"dating/internal/handler/rest"
	"dating/internal/handler/rest/mesc"
	"dating/internal/handler/rest/oauth"
	"dating/internal/handler/rest/profile"
	"dating/platform/logger"
)

type Handler struct {
	// TODO implement
	oauth   rest.OAuth
	profile rest.Profile
	mesc    rest.Mesc
}

func InitHandler(module Module, log logger.Logger) Handler {
	return Handler{
		// TODO implement
		profile: profile.InitProfile(log, module.ProfileModule),
		oauth:   oauth.InitAuth(log, module.AuthModule),
		mesc:    mesc.InitMesc(log, module.MescModule),
	}
}
