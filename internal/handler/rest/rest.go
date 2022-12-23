package rest

import (
	"github.com/gin-gonic/gin"
)

type OAuth interface {
	Test(ctx *gin.Context)
}

type Profile interface {
	Register(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	GetProfile(ctx *gin.Context)

	GetCustomers(ctx *gin.Context)
	DiscoverNewUsers(ctx *gin.Context)
}

type Mesc interface {
	CreateCountry(ctx *gin.Context)
	DeleteCountry(ctx *gin.Context)
	GetCountries(ctx *gin.Context)

	CreateState(ctx *gin.Context)
	DeleteState(ctx *gin.Context)
	GetStates(ctx *gin.Context)

	CreateEthnicity(ctx *gin.Context)
	DeleteEthnicity(ctx *gin.Context)
	UploadImage(ctx *gin.Context)
	GetEthnicities(ctx *gin.Context)
}
