package test

import (
	"github.com/stretchr/testify/assert"
	"simpledy/repository"
	"testing"
)

//var userInfo = model.UserInformation{
//	Id:            2,
//	Name:          "",
//	FollowCount:   0,
//	FollowerCount: 0,
//	User:          model.User{},
//	UserId:        6,
//}

func TestFindUserInfoByUserId(t *testing.T) {
	userInfo := repository.FindUserInfoByUserId(6)

	assert.Equal(t, int64(2), userInfo.Id)
}
