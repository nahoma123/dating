package module

import (
	"context"
	"crypto/rsa"
	"dating/internal/constant/model"

	"github.com/golang-jwt/jwt/v4"
)

type AuthModule interface {
	VerifyToken(signingMethod jwt.SigningMethod, token string, pk *rsa.PublicKey) (bool, *jwt.RegisteredClaims)
	GetUserStatus(ctx context.Context, Id string) (string, error)
}

type ProfileModule interface {
	GetUserProfile(ctx context.Context, Id string) (*model.Profile, error)
	RegisterUserProfile(ctx context.Context, profile *model.Profile) (*model.Profile, error)
	UpdateUserProfile(ctx context.Context, profile *model.Profile) (*model.Profile, error)
}
