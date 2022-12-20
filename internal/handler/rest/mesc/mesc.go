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
func (msc *mesc) CreateEthnicity(ctx *gin.Context) {
	ethnicity := &model.Ethnicity{}
	err := ctx.ShouldBind(&ethnicity)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}

	id := ctx.Param("country_id")
	ethnicity.CountryId = id
	ethnicity, err = msc.mescModule.CreateEthnicity(ctx, ethnicity)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, ethnicity, nil)
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
func (msc *mesc) DeleteEthnicity(ctx *gin.Context) {
	id := ctx.Param("ethnicity_id")
	err := msc.mescModule.DeleteEthnicity(ctx, id)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, "deleted successfully", nil)
}

// DeleteState implements rest.Mesc
func (msc *mesc) DeleteState(ctx *gin.Context) {
	id := ctx.Param("state_id")
	err := msc.mescModule.DeleteState(ctx, id)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, "deleted successfully", nil)
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
func (msc *mesc) GetEthnicities(ctx *gin.Context) {
	countryId := ctx.Param("country_id")
	ftr := constant.ParseFilterPagination(ctx)
	ftr = constant.AddFilter(*ftr, "country_id", countryId, "=")

	states, err := msc.mescModule.GetEthnicities(ctx, ftr)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, states, ftr)
}

// GetStates implements rest.Mesc
func (msc *mesc) GetStates(ctx *gin.Context) {

	countryId := ctx.Param("country_id")
	ftr := constant.ParseFilterPagination(ctx)
	ftr = constant.AddFilter(*ftr, "country_id", countryId, "=")

	states, err := msc.mescModule.GetStates(ctx, ftr)
	if err != nil {
		msc.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, states, ftr)

}

func (msc *mesc) UploadImage(ctx *gin.Context) {
	cld := constant.Credentials()
	constant.UploadImage(cld, ctx)
}
