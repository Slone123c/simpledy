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
	// videoIds, _ := repository.FindFavoriteListIdByUserId(user_id)
	favoriteList, favoriteListLength := repository.FindFavoriteListIdByUserId(int64(userId))
	favoriteVideoIdList := make([]int, favoriteListLength)
	fmt.Println(favoriteList)
	for i := 0; i < favoriteListLength; i++ {
		favoriteVideoIdList[i] = int(favoriteList[i].VideoId)
		//fmt.Println("喜欢的视频Id", favoriteVideoIdList[i])
	}

	//根据视频Id获取到视频对应的信息
	favoriteVideoList := make([]model.Video, favoriteListLength)
	for i := 0; i < favoriteListLength; i++ {
		favoriteVideoList[i], _ = repository.FindVideoInfoByVideoId(favoriteVideoIdList[i])
		//fmt.Println("喜欢的视频信息：", favoriteVideoList[i].Title)
	}

	favoriteVideoAuthorList := make([]model.UserInformation, favoriteListLength)
	for i := 0; i < favoriteListLength; i++ {
		favoriteVideoAuthorList[i] = repository.FindUserInfoByUserId(int(favoriteVideoList[i].AuthorId))
		//fmt.Println("作者是：", favoriteVideoAuthorList[i].Name)
	}

	//根据视频信息获取到对应的作者信息

	statusCode = 0
	statusMsg = "喜欢的视频列表返回成功"
	resp := model.GenerateFavoriteListResponse(int64(statusCode), statusMsg, favoriteVideoList, favoriteVideoAuthorList)
	return resp
}
