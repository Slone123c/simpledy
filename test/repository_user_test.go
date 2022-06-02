package test

import (
	"github.com/stretchr/testify/assert"
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

func TestIfUserExistsById(t *testing.T) {
	res := repository.IfUserExistsById(1)
	assert.Equal(t, true, res)
	res = repository.IfUserExistsById(2)
	assert.Equal(t, false, res)
}

func TestIfUserExistsByName(t *testing.T) {
	res := repository.IfUserExistsByUsername(username)
	assert.Equal(t, true, res)
	res = repository.IfUserExistsByUsername("unknow")
	assert.Equal(t, false, res)
}

func TestFindUserByUserName(t *testing.T) {
	res := repository.FindUserByUserName(username)
	assert.Equal(t, "testUser", res.Username)
	assert.Equal(t, "123456", res.Password)
	assert.Equal(t, uint(1), res.ID)
}

func TestInsertNewUser(t *testing.T) {
	repository.InsertNewUser(user)
	res := repository.IfUserExistsByUsername(user.Username)
	assert.Equal(t, true, res)
	repository.DeleteUserByUsername(user.Username)
	res = repository.IfUserExistsByUsername(user.Username)
	assert.Equal(t, false, res)
}

func TestDeleteUserByUsername(t *testing.T) {
	repository.InsertNewUser(user)
	res := repository.DeleteUserByUsername(user.Username)
	assert.Equal(t, true, res)
}

//func TestCustom(t *testing.T) {
//	var user = model.User{}
//	e := reflect.ValueOf(&user).Elem()
//
//	//result := db.Where(varName+" = ?", name).First(&user)
//	for i := 0; i < e.NumField(); i++ {
//		varName := e.Type().Field(i).Name
//		fmt.Println(varName)
//	}
//}
