package handler

import (
	"errors"
	"simpledy/model"
	"simpledy/repository"
)

func HandleFavoriteActionPost(token string, video_id string, action_type string) (model.FavoriteActionResponse, error) {
	var statusCode = -1
	var statusMsg = ""
	var err error
	//检测用户点赞状态是否改变
	change := repository.IfFavoriteActionYes(token, video_id, action_type)
	//点赞状态未改变，行为操作失败
	if change != true {
		statusCode = -1
		if action_type == "1" {
			err = errors.New("点赞失败")
			statusMsg = err.Error()
		} else {
			err = errors.New("取消点赞失败")
			statusMsg = err.Error()
		}
	} else {
		//点赞状态改变，设置响应消息
		statusCode = 0
		if action_type == "1" {
			statusMsg = "点赞成功"
		} else {
			statusMsg = "取消点赞成功"
		}
	}
	//更新并返回响应
	resp := generateFavoriteActionResponse(int32(statusCode), statusMsg)
	return resp, err
}

//响应整理
func generateFavoriteActionResponse(statusCode int32, statusMsg string) model.FavoriteActionResponse {
	return model.FavoriteActionResponse{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	}
}
