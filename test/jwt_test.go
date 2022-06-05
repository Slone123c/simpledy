package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"simpledy/model"
	"simpledy/utils"
	"testing"
)

//var user = model.User{
//	Name:     "zhangsan",
//	Password: "12345678",
//}
var testUser = model.User{
	Id:       1,
	Password: "123456",
	Username: "zhangsan",
}

func TestCreateToken(t *testing.T) {
	tokenString, err := utils.CreateToken(testUser)
	tokenString, err = utils.CreateToken(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)
}

func TestParseToken(t *testing.T) {
	tokenString, err := utils.CreateToken(testUser)
	tokenString, err = utils.CreateToken(user)
	if err != nil {
		panic(err)
	}
	_, claims := utils.ParseToken(tokenString)

	assert.Equal(t, claims["id"], int64(1))
	assert.Equal(t, claims["username"], "zhangsan")
}
