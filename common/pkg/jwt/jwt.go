package jwt

import (
	"errors"
	"fmt"
	"foodV5/common/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var key = []byte("ABCDEFG12334567890")

func GetTokenByUserId(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add((time.Second * 3600 * 24)).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"iss": config.C.Jwt.Key,
	})
	return token.SignedString(key)
}

func secret() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		return key, nil
	}
}

func GetUserIdByToken(tokenString string) (userId int64, err error) {
	defer func() {
		if err := recover(); err != nil {
			err = errors.New("错误的令牌")
			return
		}
	}()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)

	fmt.Println(claims)
	if !ok {
		err = errors.New("解析错误！")
	}
	if !token.Valid {
		err = errors.New("令牌错误！")
	}
	userId = int64(claims["id"].(float64))
	return
}
