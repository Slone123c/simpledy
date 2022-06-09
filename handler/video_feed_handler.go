package handler

import (
	"simpledy/model"
	"simpledy/repository"
)

func HandleVideoFeedGet(latestTime int, token string) model.VideoFeedResponse {
	var statusCode = -1
	var statusMsg = ""
	videos, videoNum := repository.FindVideosBefore(latestTime)
	authors := make([]model.UserInformation, videoNum)
	for i := 0; i < videoNum; i++ {
		video := videos[i]
		authors[i] = repository.FindUserInfoByUserId(int(video.AuthorId))
	}
	newLatestTime := videos[0].CreatedAt // 更新 latest_time
	statusCode = 0
	statusMsg = "视频返回成功"
	resp := model.GenerateVideoFeedResponse(int64(statusCode), statusMsg, newLatestTime, videos, authors)
	return resp

}
