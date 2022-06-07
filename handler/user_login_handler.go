package handler

import (
	"errors"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandleLoginPost(username string, password string) (model.UserLoginResponse, error) {

	var statusCode = -1
	var statusMsg = ""
	var userId = 0
	var token = ""
	var err error
	exist := repository.IfUserExistsByUsername(username)
	if exist {
		user := repository.FindUserByUserName(username)
		// 验证密码
		if user.Password == password {
			statusCode = 0
			statusMsg = "登陆成功"
			userId = int(user.Id)
			token, _ = utils.CreateToken(user)
		} else {
			statusCode = -2
			statusMsg = "密码错误"
			err = errors.New("用户密码输入错误")
		}
		// 检查用户是否存在
	} else {

		// 设置响应消息
		statusCode = -1
		statusMsg = "用户名不存在"
		err = errors.New("用户名不存在")
	}
	// 更新并返回响应
	resp := utils.GenerateUserLoginResponse(int32(statusCode), statusMsg, userId, token)
	return resp, err
}
