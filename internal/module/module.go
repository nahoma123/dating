package module

import (
	"context"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v4"
)

type AuthModule interface {
	VerifyToken(signingMethod jwt.SigningMethod, token string, pk *rsa.PublicKey) (bool, *jwt.RegisteredClaims)
	GetUserStatus(ctx context.Context, Id string) (string, error)
}
