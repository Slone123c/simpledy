package handler

import (
	"fmt"
	"log"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandlerRelationFollowListGet(token string,
	user_id int64) model.RelationFollowListResponseAll {
	var statusCode = -1
	var statusMsg = ""

	//将Token转换为UserId
	claims, err := utils.ParseToken(token)
	var userId int
	if err != nil {
		log.Print(err)
	}
	userId = int(claims["userId"].(float64))

	follows, _ := repository.FindFollowByUserId(userId)
	fmt.Println(follows)
	followsInfo := repository.FindFollowInfoByUserId(follows)
	fmt.Println(followsInfo)
	statusCode = 0
	statusMsg = "关注列表返回成功"
	resp := model.GenerateRelationFollowListResponse(int64(statusCode), statusMsg, followsInfo)
	fmt.Println(resp)
	return resp
}
