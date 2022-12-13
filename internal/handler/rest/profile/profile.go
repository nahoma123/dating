package profile

import (
	"dating/internal/constant"
	"dating/internal/constant/errors"
	"dating/internal/constant/model"
	"dating/internal/handler/rest"
	"dating/internal/module"
	"dating/platform/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type profile struct {
	logger        logger.Logger
	ProfileModule module.ProfileModule
}

func InitProfile(logger logger.Logger, profileModule module.ProfileModule) rest.Profile {
	return &profile{
		logger,
		profileModule,
	}
}

func (o *profile) Register(ctx *gin.Context) {
	profile := &model.Profile{}
	err := ctx.ShouldBind(&profile)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidUserInput.Wrap(err, "invalid input"))
		return
	}

	profile, err = o.ProfileModule.RegisterUserProfile(ctx, profile)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, profile, nil)
}
