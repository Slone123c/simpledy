package handler

import (
	"errors"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandleInfoGet(userId int, token string) (model.UserInfoResponse, error) {
	claims, err := utils.ParseToken(token)
	tokenUserId := claims["userId"].(float64)
	var statusCode = -1
	var statusMsg = ""
	var userInfo model.UserInformation
	// 验证 token
	if int(tokenUserId) == userId {
		statusCode = 0
		statusMsg = "验证用户成功，取得用户数据..."
		userInfo = repository.FindUserInfoByUserId(userId)
	} else {
		statusMsg = "验证用户失败"
		err = errors.New("验证用户失败")
	}
	// 更新并返回响应
	resp := model.GenerateUserInfoResponse(int32(statusCode), statusMsg, userInfo, int64(userId))
	return resp, err
}
