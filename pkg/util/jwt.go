package util

import (
	"time"

	"github.com/OswinZheng/gin-web-F/configs"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate jwt token
func GenerateToken(username, password string) (string, error) {
	jwtSecret := []byte(configs.Get().JWTConfig.Secret)
	nowTime := time.Now()
	expireTime := nowTime.Add(configs.Get().JWTConfig.EffectTime)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-seed-project",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parse jwt token
func ParseToken(token string) (*Claims, error) {
	jwtSecret := []byte(configs.Get().JWTConfig.Secret)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
