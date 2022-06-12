package handler

import (
	"fmt"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
	"time"
)

func HandleCommentPublishPost(token string, videoId int, commentText string) model.CommentActionResponse {
	statusCode := 1
	var statusMsg string
	var resp model.CommentActionResponse
	if commentText == "" {
		statusCode = -1
		statusMsg := "评论不能为空！"
		resp = model.GenerateCommentActionResponse(statusCode, statusMsg)
	} else {
		userId, err := utils.ParseTokenAndGetUserId(token)
		if err != nil {
			statusMsg = err.Error()
		} else {
			comment := model.Comment{
				Content:   commentText,
				UserId:    int64(userId),
				VideoId:   int64(videoId),
				CreatedAt: time.Now().Unix(),
			}
			repository.UpdateVideoCommentNumberPlusOneByVideoId(videoId)
			repository.InsertNewComment(comment)
			statusMsg = "评论发布成功！"
			statusCode = 0
		}
		resp = model.GenerateCommentActionResponse(statusCode, statusMsg)
	}
	return resp
}

func HandleCommentListGet(videoId int) model.CommentListResponse {
	var statusCode int
	var statusMsg string
	var resp model.CommentListResponse
	comments, commentNum := repository.FindAllCommentsByVideoId(videoId)
	if commentNum == 0 {
		fmt.Println("未找到任何评论")
		statusCode = 0
		statusMsg = "没有任何评论！"
		resp = model.GenerateCommentListResponse(statusCode, statusMsg, nil, nil, false)
	} else {
		fmt.Println("正在加载评论。。。")
		users := make([]model.UserInformation, commentNum)
		for i := 0; i < commentNum; i++ {
			users[i] = repository.FindUserInfoByUserId(int(comments[i].UserId))
		}
		statusCode = 0
		statusMsg = "评论已加载"
		resp = model.GenerateCommentListResponse(statusCode, statusMsg, comments, users, true)
	}
	return resp
}

func HandleCommentDeletePost(videoId int, commentId int) model.CommentActionResponse {
	var statusCode = 1
	var statusMsg string
	var resp model.CommentActionResponse

	res := repository.DeleteCommentByVideoIdAndCommentID(videoId, commentId)
	if res == 1 {
		statusCode = 0
		statusMsg = "评论删除成功！"
		repository.UpdateVideoCommentNumberMinusOneByVideoId(videoId)
	} else {
		statusMsg = "该评论不存在。"
	}
	resp = model.GenerateCommentActionResponse(statusCode, statusMsg)
	return resp
}
