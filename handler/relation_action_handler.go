package handler

import (
	"errors"
	"fmt"
	"log"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandlerRelationActionPost(token string, user_id int64, touser_id int64,
	action_type int) (model.FavoriteActionResponse, error) {
	var statusCode = -1
	var statusMsg = ""
	var err error

	claims, err := utils.ParseToken(token)
	var userId int64
	if err != nil {
		log.Print(err)
	}
	userId = int64(claims["userId"].(float64))
	fmt.Println(userId, touser_id)

	if userId == touser_id {
		fmt.Println("你不能关注你自己哈")
	} else {
		//检测Relation表中是否存在此关注信息
		_, exist := repository.IfRelationActionYes(userId, touser_id)
		if exist > 0 {
			//如果存在，就从Relation表中移除该数据
			repository.DeleteRelationInformation(userId, touser_id)
		} else {
			//如果不存在，就新建该关注信息到Relation表中
			newRelationInformation := model.Relation{
				UserId:   userId,
				ToUserId: touser_id,
			}
			repository.CreateRelationInformation(newRelationInformation)
		}
		change := repository.UpdateRelationCount(userId, touser_id, action_type)
		// 状态未改变，行为操作失败
		if !change {
			statusCode = -1
			if action_type == 1 {
				err = errors.New("关注失败")
				statusMsg = err.Error()
			} else {
				err = errors.New("取关失败，请重试")
				statusMsg = err.Error()
			}
		} else {
			//点赞状态改变，设置响应消息
			statusCode = 0
			if action_type == 1 {
				statusMsg = "关注成功"
			} else {
				statusMsg = "取关成功"
			}
		}
	}
	//更新并返回响应
	resp := model.GenerateFavoriteActionResponse(int32(statusCode), statusMsg)
	return resp, err
}
