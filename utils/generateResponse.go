package utils

import "simpledy/model"

func GenerateUserLoginResponse(statusCode int32, statusMsg string, userId int, token string) model.UserLoginResponse {
	return model.UserLoginResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
		UserId:     userId,
		Token:      token,
	}
}

func GenerateUserInfoResponse(statusCode int32, statusMsg string, information model.UserInformation, userId int64) model.UserInfoResponse {
	userRsp := model.UserRsp{
		ID:            userId,
		Name:          information.Name,
		FollowCount:   information.FollowCount,
		FollowerCount: information.FollowerCount,
		IsFollow:      false,
	}
	return model.UserInfoResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
		UserRsp:    userRsp,
	}
}
