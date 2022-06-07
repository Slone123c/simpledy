package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"simpledy/model"
	"simpledy/repository"
	"testing"
)

var id = 1
var username = "testUser"
var user = model.User{
	Username: "testNewUser",
	Password: "123456",
}

func TestFindUserByUserName(t *testing.T) {
	res, _ := repository.FindUserByUserName(username)
	assert.Equal(t, "testUser", res.Username)
	assert.Equal(t, "123456", res.Password)
	assert.Equal(t, uint(1), res.Id)
}

func TestInsertNewUser(t *testing.T) {
	//repository.InsertNewUser(user)
	//res := repository.IfUserExistsByUsername(user.Username)
	//assert.Equal(t, true, res)
	//repository.DeleteUserByUsername(user.Username)
	//res = repository.IfUserExistsByUsername(user.Username)
	//assert.Equal(t, false, res)
}

func TestDeleteUserByUsername(t *testing.T) {
	repository.InsertNewUser(user)
	res := repository.DeleteUserByUsername(user.Username)
	assert.Equal(t, true, res)
}

//func TestCustom1(t *testing.T) {
//	var user = model.User{}
//	e := reflect.ValueOf(&user).Elem()
//
//	//result := db.Where(varName+" = ?", name).First(&user)
//	for i := 0; i < e.NumField(); i++ {
//		varName := e.Type().Field(i).Name
//		fmt.Println(varName)
//	}
//}

func TestCustom(t *testing.T) {
	var user = model.UserInformation{}
	e := reflect.ValueOf(&user).Elem()
	//result := db.Where(varName+" = ?", name).First(&user)
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		fmt.Println(varName)
	}
}
