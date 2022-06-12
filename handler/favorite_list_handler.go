package handler

import (
	"fmt"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandleFavoriteListGet(user_id int64, token string) model.FavoriteListResponse {
	var statusCode = -1
	var statusMsg = ""
	userId, _ := utils.ParseTokenAndGetUserId(token)

	//获取用户喜欢的所有Video的Id
	// videoIds, _ := repository.FindVideoIdByUserId(user_id)
	videoIds, _ := repository.FindVideoIdByUserId(int64(userId))
	fmt.Println(videoIds)

	//根据视频Id获取到视频对应的信息
	videos, _ := repository.FindVideoInfoByVideoId(videoIds)
	fmt.Println(videos)

	//根据视频信息获取到对应的作者信息
	authors, _ := repository.FindAuthorInfoByVideoInfo(videos)
	fmt.Println(authors)

	statusCode = 0
	statusMsg = "喜欢的视频列表返回成功"
	resp := model.GenerateFavoriteListResponse(int64(statusCode), statusMsg, videos, authors)
	return resp
}
