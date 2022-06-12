package service

// import (
// 	"log"
// 	"net/http"
// 	"simpledy/handler"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func RelationAction(c *gin.Context) {
// 	//获取用户信息
// 	token := c.PostForm("token")
// 	//获取当前操作用户ID
// 	user_idStr := c.PostForm("user_id")
// 	user_id, _ := strconv.Atoi(user_idStr)
// 	//获取被关注用户ID
// 	toUser_idStr := c.PostForm("user_id")
// 	toUser_id, _ := strconv.Atoi(toUser_idStr)
// 	//获取用户行为
// 	action_type := c.PostForm("action_type")

// 	resp, err := handler.HandlerRelationActionPost(token, int64(user_id), int64(toUser_id), action_type)
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	//返回响应信息
// 	c.JSON(http.StatusOK, resp)
// }

// func RelationFollowList(c *gin.Context) {
// 	//获取请求参数
// 	user_idStr := c.Query("user_id")
// 	user_id, _ := strconv.Atoi(user_idStr)
// 	token := c.Query("token")

// 	resp := handler.HandlerRelationFollowListGet(token, int64(user_id))

// 	//返回响应信息
// 	c.JSON(http.StatusOK, resp)
// }

// func RelationFollowerList(c *gin.Context) {
// 	//获取请求参数
// 	user_idStr := c.Query("user_id")
// 	user_id, _ := strconv.Atoi(user_idStr)
// 	token := c.Query("token")

// 	resp := handler.HandlerRelationFollowerListGet(token, int64(user_id))

// 	//返回响应信息
// 	c.JSON(http.StatusOK, resp)
// }
