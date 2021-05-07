package utils

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"lohon.cm/msvc/auth/db"
)

var (
	_JWTPrivateKeyPath = "/var/lib/lohon/jwt-keys/jwtRS256.key"
	_JWTPublicKeyPath  = "/var/lib/lohon/jwt-keys/jwtRS256.pub"
)

type CustomClaims struct {
	*jwt.StandardClaims
	db.User
}

func GetJWT(user db.User) (string, error) {
	signBytes, err := ioutil.ReadFile(_JWTPrivateKeyPath)
	if err != nil {
		return "", err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()
	ttl := 1 * time.Hour

	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims = &CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: now.Add(ttl).Unix(),
			IssuedAt:  now.Unix(),
		},
		user,
	}

	return token.SignedString(signKey)
}

func VerifyJWT(token string) (interface{}, error) {
	verifyBytes, err := ioutil.ReadFile(_JWTPublicKeyPath)
	if err != nil {
		return "", err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return "", fmt.Errorf("validate: parse key: %w", err)
	}

	tokenParsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}

		return verifyKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if !ok || !tokenParsed.Valid {
		return nil, fmt.Errorf("validated: invalid")
	}

	return claims["dat"], nil
}
