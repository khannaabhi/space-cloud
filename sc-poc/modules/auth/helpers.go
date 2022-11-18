package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	"github.com/spacecloud-io/space-cloud/modules/auth/types"
)

func getSigninMethod(alg types.JWTAlg) jwt.SigningMethod {
	switch alg {
	case types.HS256:
		return jwt.SigningMethodHS256
	case types.RS256, types.RS256Public, types.JwkURL:
		return jwt.SigningMethodRS256
	}

	return jwt.SigningMethodHS256
}

func getSigningSecret(secret *types.AuthSecret) interface{} {
	if secret.Alg == types.HS256 {
		return []byte(secret.Value)
	}

	// TODO: Return private key here
	return nil
}

func getVerifyingSecret(secret *types.AuthSecret) interface{} {
	if secret.Alg == types.HS256 {
		return []byte(secret.Value)
	}

	// TODO: Return public key here
	return nil
}

func (a *App) getPrimarySecret() *types.AuthSecret {
	for _, s := range a.secrets {
		if s.IsPrimary {
			return s
		}
	}

	return a.secrets[0]
}

func (a *App) getSecretFromKID(kid string) (*types.AuthSecret, bool) {
	for _, s := range a.secrets {
		if s.KID == kid {
			return s, true
		}
	}

	return nil, false
}

func verifyClaims(claims interface{}, configValues []string) error {
	switch claimFromToken := claims.(type) {
	case string:
		for _, cmp := range configValues {
			if claimFromToken == cmp {
				return nil
			}
		}
	case []interface{}:
		for _, audClaim := range claimFromToken {
			for _, c := range configValues {
				if audClaim == c {
					return nil
				}
			}
		}
	default:
		return errors.New("invalid type provided for claim in jwt token")
	}
	return errors.New("unable to verify claim of jwt token")
}

func getTokenFromHeader(r *http.Request) (string, bool) {
	// Get the JWT token from header
	tokens, ok := r.Header["Authorization"]
	if !ok {
		tokens = []string{""}
	}

	arr := strings.Split(tokens[0], " ")
	if strings.ToLower(arr[0]) == "bearer" && len(arr) >= 2 {
		return arr[1], true
	}

	return "", false
}