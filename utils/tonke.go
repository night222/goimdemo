package utils

import (
	"errors"
	"goimdemo/common"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

// token的加密和解密
func GenerateToken(id uint, name string) (string, error) {
	claims := TokenClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(common.ConfigData.ExpiredAt) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   common.ConfigData.Subject,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(common.ConfigData.Singin)) //这里注意一定一定要用byte切片不然无效
}

// token解密
func ParseToken(tokenStr string) (TokenClaims, error) {
	iTokenClaims := TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iTokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.ConfigData.Singin), nil
	})
	if err != nil && !token.Valid {
		err = errors.New("Invail Token")
	}
	return iTokenClaims, err
}
