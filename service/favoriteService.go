package service

import (
	"log"
	"net/http"
	"simpledy/handler"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	//获取用户信息
	token := c.PostForm("token")
	//获取当前视频
	video_id := c.PostForm("video_id")
	//获取用户行为
	action_type := c.PostForm("action_type")

	resp, err := handler.HandleFavoriteActionPost(token, video_id, action_type)
	if err != nil {
		log.Print(err)
	}

	//返回响应信息
	c.JSON(http.StatusOK, resp)
}

// func FavoriteList(c *gin.Context) {
// 	//获取请求参数
// 	user_id := c.Query("user_id")
// 	token := c.Query("token")

// 	resp, err := handler.HandleFavoriteListPost(user_id, token)
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	//返回响应信息
// 	c.JSON(http.StatusOK, resp)
// }
