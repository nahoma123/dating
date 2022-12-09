package middleware

import (
	"context"
	"crypto/rsa"
	"dating/internal/constant"
	"dating/internal/constant/errors"
	"dating/internal/module"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware interface {
	Authentication() gin.HandlerFunc
}

type authMiddleware struct {
	auth         module.AuthModule
	ssoPublicKey *rsa.PublicKey
}

func InitAuthMiddleware(
	auth module.AuthModule, ssoPublicKey *rsa.PublicKey) AuthMiddleware {
	return &authMiddleware{
		auth,
		ssoPublicKey,
	}
}

func (a *authMiddleware) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			Err := errors.ErrInvalidToken.New("Unauthorized")
			ctx.Error(Err)
			ctx.Abort()
			return
		}

		tokenString := authHeader[len(bearer):]
		valid, claims := a.auth.VerifyToken(jwt.SigningMethodPS512, tokenString, a.ssoPublicKey)
		if !valid {
			Err := errors.ErrAuthError.New("Unauthorized")
			ctx.Error(Err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userStatus, err := a.auth.GetUserStatus(ctx.Request.Context(), claims.Subject)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		if userStatus != constant.Active {
			Err := errors.ErrAuthError.Wrap(nil, "Your account has been deactivated, Please activate your account.")
			ctx.Error(Err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), constant.Context("x-user-id"), claims.Subject))
		ctx.Next()
	}
}
