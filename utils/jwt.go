package utils

import (
	"time"

	"github.com/panda8z/time-book/pkg/settings"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(settings.JwtSecret)

// Claims for access token
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken create a token with given params
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	// token validly for 3 hours
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "time-book",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(jwtSecret)

	return token, err
}

// ParseToken can turn a token string to Claims object
func ParseToken(token string) (*Claims, error) {
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
