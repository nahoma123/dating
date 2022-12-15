package mesc

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

type mesc struct {
	logger     logger.Logger
	mescModule module.MescModule
}

func InitMesc(logger logger.Logger, mescModule module.MescModule) rest.Mesc {
	return &mesc{
		logger,
		mescModule,
	}
}

// CreateCountry implements rest.Mesc
func (msc *mesc) CreateCountry(ctx *gin.Context) {
	mesc := &model.Country{}
	err := ctx.ShouldBind(&mesc)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidUserInput.Wrap(err, "invalid input"))
		return
	}

	mesc, err = msc.mescModule.CreateCountry(ctx, mesc)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, msc, nil)

}

// CreateEthnicity implements rest.Mesc
func (*mesc) CreateEthnicity(ctx *gin.Context) {
	panic("unimplemented")
}

// CreateState implements rest.Mesc
func (*mesc) CreateState(ctx *gin.Context) {
	panic("unimplemented")
}

// DeleteCountry implements rest.Mesc
func (*mesc) DeleteCountry(ctx *gin.Context) {
	panic("unimplemented")
}

// DeleteEthnicity implements rest.Mesc
func (*mesc) DeleteEthnicity(ctx *gin.Context) {
	panic("unimplemented")
}

// DeleteState implements rest.Mesc
func (*mesc) DeleteState(ctx *gin.Context) {
	panic("unimplemented")
}

// GetCountries implements rest.Mesc
func (*mesc) GetCountries(ctx *gin.Context) {
	panic("unimplemented")
}

// GetEthnicities implements rest.Mesc
func (*mesc) GetEthnicities(ctx *gin.Context) {
	panic("unimplemented")
}

// GetStates implements rest.Mesc
func (*mesc) GetStates(ctx *gin.Context) {
	panic("unimplemented")
}
