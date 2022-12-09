package oauth

import (
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func (o *oauth) VerifyToken(signingMethod jwt.SigningMethod, token string, pk *rsa.PublicKey) (bool, *jwt.RegisteredClaims) {
	claims := &jwt.RegisteredClaims{}

	segments := strings.Split(token, ".")
	if len(segments) < 3 {
		return false, claims
	}
	err := signingMethod.Verify(strings.Join(segments[:2], "."), segments[2], pk)
	if err != nil {
		return false, claims
	}

	if _, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSAPSS); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return pk, nil
	}); err != nil {
		return false, claims
	}
	return true, claims
}
