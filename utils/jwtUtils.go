package utils

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
		"userId":      user.Id,
		"username":    user.Username,
		"expiredTime": expirationTime,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtKey)

	fmt.Println(tokenString, err)
	return tokenString, err
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if token.Valid {
		fmt.Println("令牌有效。")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("无效令牌")
		err = errors.New("无效令牌")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("令牌已过期")
		err = errors.New("令牌已过期")
	} else {
		fmt.Println("无法处理令牌", err)
		err = errors.New("无法处理令牌")
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return claims, err
	} else {
		return nil, nil
	}

}

func AuthUser(userId int, token string) (bool, string) {
	ok := false
	returnMsg := ""
	tokenUserId, err := ParseTokenAndGetUserId(token)
	if err != nil {
		fmt.Println("1：")
		return ok, err.Error()
	}
	// 验证 token
	fmt.Println("tokenid = ", tokenUserId)
	fmt.Println("userid = ", userId)
	if tokenUserId == userId {
		returnMsg = "验证用户成功"
		ok = true
	} else {
		fmt.Println("3：")
		returnMsg = "验证用户失败"
	}
	return ok, returnMsg
}

func ParseTokenAndGetUserId(tokenString string) (int, error) {
	claims, err := ParseToken(tokenString)
	userId := claims["userId"].(float64)
	return int(userId), err
}
