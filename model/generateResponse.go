package model

func GenerateUserLoginResponse(statusCode int32, statusMsg string, userId int, token string) UserLoginResponse {
	return UserLoginResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
		UserId:     userId,
		Token:      token,
	}
}

func GenerateUserInfoResponse(statusCode int32, statusMsg string, information UserInformation, userId int64) UserInfoResponse {
	userRsp := UserRsp{
		ID:            userId,
		Name:          information.Name,
		FollowCount:   information.FollowCount,
		FollowerCount: information.FollowerCount,
		IsFollow:      false,
	}
	return UserInfoResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
		UserRsp:    userRsp,
	}
}

func GenerateVideoFeedResponse() {

}
