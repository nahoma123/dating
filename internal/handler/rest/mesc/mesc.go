package mesc

import (
	"dating/internal/constant"
	"dating/internal/constant/errors"
	"dating/internal/constant/model"
	"dating/internal/handler/rest"
	"dating/internal/module"
	"dating/platform/logger"
	"net/http"
	"strconv"

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
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
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
func (msc *mesc) CreateState(ctx *gin.Context) {
	state := &model.State{}
	err := ctx.ShouldBind(&state)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	id := ctx.Param("country_id")
	state.CountryId = id
	state, err = msc.mescModule.CreateState(ctx, state)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, state, nil)
}

// DeleteCountry implements rest.Mesc
func (msc *mesc) DeleteCountry(ctx *gin.Context) {
	id := ctx.Param("country_id")
	err := msc.mescModule.DeleteCountry(ctx, id)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, "deleted successfully", nil)
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
func (msc *mesc) GetCountries(ctx *gin.Context) {
	page := ctx.Query("page")
	perPage := ctx.Query("per_page")
	if page == "" {
		page = "1"
	}
	if perPage == "" {
		perPage = "10"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	perPageInt, err := strconv.Atoi(perPage)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	countries, metaData, err := msc.mescModule.GetCountries(ctx, pageInt, perPageInt)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, countries, metaData)

}

// GetEthnicities implements rest.Mesc
func (*mesc) GetEthnicities(ctx *gin.Context) {
	panic("unimplemented")
}

// GetStates implements rest.Mesc
func (*mesc) GetStates(ctx *gin.Context) {
	panic("unimplemented")
}
