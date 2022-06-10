package test

import (
	"github.com/stretchr/testify/assert"
	"simpledy/handler"
	"testing"
)

/*
type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int    `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}
*/
var name = "testUser"
var pwd = "123456"

func TestHandleRegisterPost(t *testing.T) {
	res, _ := handler.HandleRegisterPost(name, pwd)
	assert.Equal(t, int32(-1), res.StatusCode)
	name = "newUser"
	res, _ = handler.HandleRegisterPost(name, pwd)
	assert.Equal(t, int32(1), res.StatusCode)
}

//func TestHandlePublishPost(t *testing.T)  {
//	res := handler.HandlePublishPost()
//}
