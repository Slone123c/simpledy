package handler

import (
	"errors"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandleRegisterPost(username string, password string) (model.UserLoginResponse, error) {
	// 检查用户是否重复
	var statusCode = -1
	var statusMsg = ""
	var userId = 0
	var token = ""
	var err error
	_, exist := repository.FindUserByUserName(username)
	if exist > 0 {
		statusCode = 1
		err = errors.New("用户名已存在")
		statusMsg = err.Error()
	} else {
		// 设置响应消息
		statusCode = 0
		statusMsg = "用户注册成功"
		// 创建新用户
		user := model.User{
			Username: username,
			Password: password,
		}

		// 数据库插入新用户
		repository.InsertNewUser(user)
		newUser, _ := repository.FindUserByUserName(username)
		userId = int(newUser.Id)
		token, _ = utils.CreateToken(newUser)
		// 创建新用户表
		userInfo := model.UserInformation{
			Name:          newUser.Username,
			FollowCount:   0,
			FollowerCount: 0,
			UserId:        newUser.Id,
		}
		repository.InsertNewUserInformation(userInfo)
	}
	// 更新并返回响应
	resp := utils.GenerateUserLoginResponse(int32(statusCode), statusMsg, userId, token)
	return resp, err
}
