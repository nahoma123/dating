package profile

import (
	"dating/internal/constant"
	"dating/internal/constant/errors"
	"dating/internal/constant/model"
	"dating/internal/handler/rest"
	"dating/internal/module"
	"dating/platform/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
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

func (o *profile) UpdateProfile(ctx *gin.Context) {
	profile := &model.Profile{}
	err := ctx.ShouldBind(&profile)
	id := ctx.Param("id")
	if err != nil || id == "" {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(errors.ErrInvalidInput.Wrap(err, "invalid input"))
		return
	}
	profile.ProfileID = id
	profile, err = o.ProfileModule.UpdateUserProfile(ctx, profile)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, profile, nil)
}

func (o *profile) GetProfile(ctx *gin.Context) {
	id := ctx.Param("id")
	profile, err := o.ProfileModule.GetUserProfile(ctx, id)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusCreated, profile, nil)
}

func (o *profile) GetCustomers(ctx *gin.Context) {
	ftr := constant.ParseFilterPagination(ctx)
	user_id := ctx.GetString("x-user-id")

	// url := ctx.Request.URL.Path
	ftr = constant.AddFilter(*ftr, "profile_id", user_id, "!=")
	customers, err := o.ProfileModule.GetCustomers(ctx, ftr)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, customers, ftr)

}

func (o *profile) DiscoverNewUsers(ctx *gin.Context) {
	ftr := constant.ParseFilterPagination(ctx)
	user_id := ctx.GetString("x-user-id")

	// url := ctx.Request.URL.Path
	ftr = constant.AddFilter(*ftr, "profile_id", user_id, "!=")

	constant.DeleteFilter(ftr, "created_at")

	oneMonthAgo := time.Now().AddDate(viper.GetInt("matching.new_users_time.year"), viper.GetInt("matching.new_users_time.month"), viper.GetInt("matching.new_users_time.month"))

	// Format the time as an ISO 8601 formatted string.
	isoString := oneMonthAgo.Format("2006-01-02T15:04:05")

	ftr = constant.AddFilter(*ftr, "created_at", isoString, "gte")
	ftr = constant.AddFilter(*ftr, "distance", fmt.Sprintf("%d", viper.GetInt("matching.nearby_distance")), "gte")
	customers, err := o.ProfileModule.GetCustomers(ctx, ftr)
	if err != nil {
		o.logger.Info(ctx, zap.Error(err).String)
		_ = ctx.Error(err)
		return
	}

	constant.SuccessResponse(ctx, http.StatusOK, customers, ftr)
}
