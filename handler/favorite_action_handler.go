package handler

import (
	"log"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
)

func HandleFavoriteActionPost(user_id int, token string, video_id int64,
	action_type int) (model.FavoriteActionResponse, error) {
	var statusCode = -1
	var statusMsg = ""
	var err error

	// //将Token转化为UserId
	claims, err := utils.ParseToken(token)
	var userId int64
	if err != nil {
		log.Print(err)
	}
	userId = int64(claims["userId"].(float64))

	//检测Favorite表中是否存在此点赞消息
	_, exist := repository.IfFavoriteActionYes(int64(userId), video_id, action_type)
	if exist > 0 && action_type == 2 {
		//如果存在，就从Favorite表中移除该数据
		repository.DeleteFavoriteInformation(int64(userId), video_id)
		repository.UpdateVideoFavoriteNumberMinusOneByVideoId(int(video_id))
		statusCode = 0
		statusMsg = "取消点赞成功！"
		resp := model.GenerateFavoriteActionResponse(int32(statusCode), statusMsg)
		return resp, err
	} else if exist > 0 {
		statusCode = 0
		statusMsg = "已经点赞过了！"
		resp := model.GenerateFavoriteActionResponse(int32(statusCode), statusMsg)
		return resp, err
	} else if exist == 0 && action_type == 1 {
		//如果不存在，就新建该点赞信息进到Favorite表
		newFavoriteInformation := model.Favorite{
			UserId:  int64(userId),
			VideoId: video_id,
		}
		statusMsg = "点赞成功！"
		statusCode = 0
		repository.CreateFavoriteInformation(newFavoriteInformation)
		repository.UpdateVideoFavoriteNumberPlusOneByVideoId(int(video_id))
		resp := model.GenerateFavoriteActionResponse(int32(statusCode), statusMsg)
		return resp, err
	} else {
		statusMsg = "点赞失败"
		statusCode = 1
		resp := model.GenerateFavoriteActionResponse(int32(statusCode), statusMsg)
		return resp, err
	}

}
