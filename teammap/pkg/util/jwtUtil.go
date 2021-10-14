package util

import (
	"teammap/pkg/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId int `json:"uid"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(userId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(30 * 24 * time.Hour) // 30天后过期

	claims := Claims{
		Atoi(userId),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(setting.AppSetting.JwtSecret))

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.AppSetting.JwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
