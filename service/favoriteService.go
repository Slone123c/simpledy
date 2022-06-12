package service

import (
	"log"
	"net/http"
	"simpledy/handler"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	//获取用户信息
	user_idStr := c.Query("user_id")
	user_id, _ := strconv.Atoi(user_idStr)
	token := c.Query("token")
	//获取当前视频
	video_idStr := c.Query("video_id")
	video_id, _ := strconv.Atoi(video_idStr)
	//获取用户行为
	action_typeStr := c.Query("action_type")
	action_type, _ := strconv.Atoi(action_typeStr)

	resp, err := handler.HandleFavoriteActionPost(user_id, token, int64(video_id), action_type)
	if err != nil {
		log.Print(err)
	}

	//返回响应信息
	c.JSON(http.StatusOK, resp)
}

func FavoriteList(c *gin.Context) {
	//获取请求参数
	user_idStr := c.Query("user_id")
	user_id, _ := strconv.Atoi(user_idStr)
	token := c.Query("token")

	resp := handler.HandleFavoriteListGet(int64(user_id), token)

	//返回响应信息
	c.JSON(http.StatusOK, resp)
}
