package controller

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"simpledy/model"
	"time"
)

var jwtKey = []byte("user_key")

func CreateToken(user model.User) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	duration, _ := time.ParseDuration("24h")
	expirationTime := time.Now().Add(duration)

	// 设置加密算法 和 Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 定义 Claims 结构体
		"id":          user.ID,
		"username":    user.Name,
		"expiredTime": expirationTime,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtKey)

	fmt.Println(tokenString, err)
	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, jwt.MapClaims) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if token.Valid {
		fmt.Println("令牌有效。")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("无效令牌。")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("令牌已过期。")
	} else {
		fmt.Println("无法处理令牌：", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return token, claims
	} else {
		return nil, nil
	}

}
