package handler

import (
	"log"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandlePublishListGet(token string) model.PublishListResponse {
	//videos, videoNum := repository.FindVideosBefore(latestTime)
	//authors := make([]model.UserInformation, videoNum)
	//for i := 0; i < videoNum; i++ {
	//	video := videos[i]
	//	authors[i] = repository.FindUserInfoByUserId(int(video.AuthorId))
	//}
	//newLatestTime := videos[0].CreatedAt // 更新 latest_time
	//resp := model.GenerateVideoFeedResponse(0, "测试", newLatestTime, videos, authors)
	var statusCode = -1
	var statusMsg = ""
	claims, err := utils.ParseToken(token)
	var userId int
	//fmt.Println("userid =", claims["userId"])
	//fmt.Println("useridInt =", claims["userId"].(int))
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
