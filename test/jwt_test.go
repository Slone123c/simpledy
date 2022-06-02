package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"simpledy/model"
	"simpledy/utils"
	"testing"
)

//var user = model.User{
//	Name:     "zhangsan",
//	Password: "12345678",
//}
var testUser = model.User{
	Model:    gorm.Model{ID: 1},
	Username: "zhangsan",
}

func TestCreateToken(t *testing.T) {
<<<<<<< HEAD
	tokenString, err := utils.CreateToken(testUser)
=======
	tokenString, err := utils.CreateToken(user)
>>>>>>> main
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)
}

func TestParseToken(t *testing.T) {
<<<<<<< HEAD
	tokenString, err := utils.CreateToken(testUser)
=======
	tokenString, err := utils.CreateToken(user)
>>>>>>> main
	if err != nil {
		panic(err)
	}
	_, claims := utils.ParseToken(tokenString)

	assert.Equal(t, claims["id"], float64(1))
	assert.Equal(t, claims["username"], "zhangsan")
}
