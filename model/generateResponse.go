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

func GenerateVideoFeedResponse(statusCode int64, statusMsg string, nextTime int64, videos []Video, authors []UserInformation) VideoFeedResponse {
	videoList := make([]VideoList, len(videos))
	authorList := make([]author, len(authors))
	for i := 0; i < len(videos); i++ {
		videoList[i].ID = videos[i].Id
		videoList[i].PlayURL = videos[i].PlayUrl
		videoList[i].CoverURL = videos[i].CoverUrl
		videoList[i].FavoriteCount = videos[i].FavoriteCount
		videoList[i].CommentCount = videos[i].CommentCount
		videoList[i].IsFavorite = videos[i].IsFavorite
		videoList[i].Title = videos[i].Title
		authorList[i].ID = authors[i].Id
		authorList[i].Name = authors[i].Name
		authorList[i].FollowCount = authors[i].FollowCount
		authorList[i].FollowerCount = authors[i].FollowerCount
		videoList[i].Author = authorList[i]
		//authorList[i].IsFollow = authors[i].
	}

	return VideoFeedResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
		NextTime:   nextTime,
		VideoList:  videoList,
	}
}
