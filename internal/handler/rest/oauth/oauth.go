package oauth

import (
	"dating/internal/constant"
	"dating/internal/constant/errors"
	"dating/internal/constant/model/dto"
	"dating/internal/handler/rest"
	"dating/internal/module"
	"dating/platform/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type auth struct {
	logger     logger.Logger
	AuthModule module.AuthModule
}

func InitAuth(logger logger.Logger, AuthModule module.AuthModule) rest.OAuth {
	return &auth{
		logger,
		AuthModule,
	}
}

func (o *auth) Test(ctx *gin.Context) {
	userParam := dto.User{}
	err := ctx.ShouldBind(&userParam)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidUserInput.Wrap(err, "invalid input"))
		return
	}

	o.logger.Info(ctx, "testing log dating")
	constant.SuccessResponse(ctx, http.StatusCreated, nil, nil)
}
