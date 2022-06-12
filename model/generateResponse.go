package model

import (
	"time"
)

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
	authorList := make([]Author, len(authors))
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

func GeneratePulishListResponse(statusCode int64, statusMsg string, videos []Video, authorInfo UserInformation) PublishListResponse {
	videoList := make([]VideoList, len(videos))
	author := Author{
		FollowCount:   authorInfo.FollowerCount,
		FollowerCount: authorInfo.FollowerCount,
		ID:            authorInfo.Id,
		IsFollow:      false,
		Name:          authorInfo.Name,
	}
	for i := 0; i < len(videos); i++ {
		videoList[i].ID = videos[i].Id
		videoList[i].PlayURL = videos[i].PlayUrl
		videoList[i].CoverURL = videos[i].CoverUrl
		videoList[i].FavoriteCount = videos[i].FavoriteCount
		videoList[i].CommentCount = videos[i].CommentCount
		videoList[i].IsFavorite = videos[i].IsFavorite
		videoList[i].Title = videos[i].Title
		videoList[i].Author = author
	}
	return PublishListResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
		VideoList:  videoList,
	}
}

func GenerateVideoPublishResponse(statusCode int64, statusMsg string) VideoPublishResponse {
	return VideoPublishResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}
func GenerateFavoriteActionResponse(statusCode int32, statusMsg string) FavoriteActionResponse {
	return FavoriteActionResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}

func GenerateFavoriteListResponse(statusCode int64, statusMsg string, videos []Video, authorInfos []UserInformation) FavoriteListResponse {
	authorList := make([]Author, len(authorInfos))
	videoList := make([]VideoList, len(videos))
	for j := 0; j < len(authorInfos); j++ {
		authorList[j].FollowCount = authorInfos[j].FollowerCount
		authorList[j].FollowerCount = authorInfos[j].FollowerCount
		authorList[j].ID = authorInfos[j].Id
		authorList[j].IsFollow = false
		authorList[j].Name = authorInfos[j].Name
	}
	for i := 0; i < len(videos); i++ {
		videoList[i].ID = videos[i].Id
		videoList[i].PlayURL = videos[i].PlayUrl
		videoList[i].CoverURL = videos[i].CoverUrl
		videoList[i].FavoriteCount = videos[i].FavoriteCount
		videoList[i].CommentCount = videos[i].CommentCount
		videoList[i].IsFavorite = videos[i].IsFavorite
		videoList[i].Title = videos[i].Title
		videoList[i].Author = authorList[i]
	}
	return FavoriteListResponse{
		StatusCode: int32(statusCode),
		StatusMsg:  statusMsg,
		VideoList:  videoList,
	}
}

func GenerateCommentActionResponse(StatusCode int, StatusMsg string) CommentActionResponse {
	return CommentActionResponse{
		StatusCode: StatusCode,
		StatusMsg:  StatusMsg,
	}
}

func GenerateCommentListResponse(statusCode int, statusMsg string, comments []Comment, users []UserInformation, hasComment bool) CommentListResponse {
	if hasComment {
		commentList := make([]CommentRespElement, len(comments))
		commentAuthorList := make([]CommentAuthorInfo, len(comments))
		for i := 0; i < len(comments); i++ {
			commentAuthorList[i].ID = users[i].UserId
			commentAuthorList[i].Name = users[i].Name
			commentAuthorList[i].FollowCount = users[i].FollowCount
			commentAuthorList[i].FollowerCount = users[i].FollowerCount
			//commentAuthorList[i].IsFollow =  users[i].
		}

		for i := 0; i < len(comments); i++ {
			commentList[i].ID = int(comments[i].Id)
			commentList[i].User = commentAuthorList[i]
			commentList[i].Content = comments[i].Content
			timeType := time.Unix(comments[i].CreatedAt, 0)
			timeString := timeType.Format("01-02")
			commentList[i].CreateDate = timeString

		}
		return CommentListResponse{
			StatusCode:  statusCode,
			StatusMsg:   statusMsg,
			CommentList: commentList,
		}
	} else {
		return CommentListResponse{
			StatusCode: statusCode,
			StatusMsg:  statusMsg,
		}
	}

}
