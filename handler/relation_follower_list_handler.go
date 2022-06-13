package handler

import (
	"fmt"
	"log"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandlerRelationFollowerListGet(token string,
	user_id int64) model.RelationFollowerListResponseAll {
	var statusCode = -1
	var statusMsg = ""

	//将Token转换为UserId
	claims, err := utils.ParseToken(token)
	var userId int
	if err != nil {
		log.Print(err)
	}
	userId = int(claims["userId"].(float64))

	followers, _ := repository.FindFollowerByUserId(userId)
	fmt.Println(followers)
	followersInfo := repository.FindFollowerInfoByUserId(followers)
	fmt.Println(followersInfo)
	statusCode = 0
	statusMsg = "粉丝列表返回成功"

	resp := repository.GenerateRelationFollowerListResponse(int64(statusCode),
		statusMsg, followersInfo, int64(userId))
	fmt.Println(resp)
	return resp
}
