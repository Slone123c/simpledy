package handler

import (
	"log"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandlePublishListGet(token string) model.PublishListResponse {
	var statusCode = -1
	var statusMsg = ""
	claims, err := utils.ParseToken(token)
	var userId int
	if err != nil {
		log.Print(err)
	}
	userId = int(claims["userId"].(float64))
	videos, _ := repository.FindVideosByUserId(userId)
	author := repository.FindUserInfoByUserId(userId)
	statusCode = 0
	statusMsg = "视频列表返回成功"
	resp := model.GeneratePulishListResponse(int64(statusCode), statusMsg, videos, author)
	return resp

}
