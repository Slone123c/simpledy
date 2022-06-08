package handler

import (
	"simpledy/model"
	"simpledy/repository"
)

func HandleVideoFeedGet(latestTime int, token string) model.VideoFeedResponse {

	videos, videoNum := repository.FindVideosBefore(latestTime)
	authors := make([]model.UserInformation, videoNum)
	for i := 0; i < videoNum; i++ {
		video := videos[i]
		authors[i] = repository.FindUserInfoByUserId(int(video.AuthorId))
	}
	newLatestTime := videos[0].CreatedAt // 更新 latest_time
	resp := model.GenerateVideoFeedResponse(0, "测试", newLatestTime, videos, authors)
	return resp

}
