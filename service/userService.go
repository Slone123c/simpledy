package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simpledy/handler"
	"simpledy/model"
	"strconv"
)

func Register(c *gin.Context) {
	fmt.Println("Registering....")
	// 获取用户输入
	username := c.Query("username")
	password := c.Query("password")

	resp, err := handler.HandleRegisterPost(username, password)
	if err != nil {
		log.Print(err)
	}
	// 返回响应信息
	c.JSON(http.StatusOK, resp)
}

func Login(c *gin.Context) {
	fmt.Println("Login....")
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println("uname:", username)
	fmt.Println("upwd:", password)
	resp, err := handler.HandleLoginPost(username, password)
	if err != nil {
		log.Print(err)
	}
	c.JSON(http.StatusOK, resp)
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	id_string := c.Query("user_id")
	id, _ := strconv.Atoi(id_string)
	//fmt.Println("用户登录token如下：", token)
	//fmt.Println("用户登录id：", id)
	resp, err := handler.HandleInfoGet(id, token)
	if err != nil {
		log.Print(err)
	}
	c.JSON(http.StatusOK, resp)
	//c.JSON(http.StatusOK, UserResponse{
	//	Response: Response{StatusCode: 0},
	//	User:     testU,
	//})

}

func Feed(c *gin.Context) {
	latestTime, _ := strconv.Atoi(c.Query("latest_time"))
	token := c.Query("token")
	resp := handler.HandleVideoFeedGet(latestTime, token)
	//if err != nil {
	//	log.Print(err)
	//}
	c.JSON(http.StatusOK, resp)

}

func PublishList(c *gin.Context) {
	token := c.Query("token")
	resp := handler.HandlePublishListGet(token)
	c.JSON(http.StatusOK, resp)
}

func Publish(c *gin.Context) {
	data, _ := c.FormFile("data")
	token := c.PostForm("token")
	title := c.PostForm("title")
	if err := c.SaveUploadedFile(data, "./static_data/videos/"+data.Filename); err != nil {
		fmt.Println(err)
	}
	//
	resp := handler.HandlePublishPost(data.Filename, token, title)
	fmt.Println("file uploaded:", data.Filename)
	fmt.Println("This is token:", token)
	fmt.Println("This is title:", title)
	c.JSON(http.StatusOK, resp)
}

func CommentAction(c *gin.Context) {
	var resp model.CommentActionResponse
	id_string := c.Query("user_id")
	token := c.Query("token")
	video_string := c.Query("video_id")
	action_type := c.Query("action_type")
	commentText := c.Query("comment_text")
	comment_id := c.Query("comment_id")
	userId, _ := strconv.Atoi(id_string)
	videoId, _ := strconv.Atoi(video_string)
	actionType, _ := strconv.Atoi(action_type)
	commentId, _ := strconv.Atoi(comment_id)
	fmt.Println("userid为：", userId)
	//fmt.Println(token)
	//fmt.Println(video_id)
	//fmt.Println(action_type)
	//fmt.Println(comment_text)
	//fmt.Println(comment_id)
	if actionType == 1 {
		fmt.Println("执行发布评论。。。")
		resp = handler.HandleCommentPublishPost(token, videoId, commentText)
	} else if actionType == 2 {
		fmt.Println("执行删除评论")
		resp = handler.HandleCommentDeletePost(videoId, commentId)
	}
	c.JSON(http.StatusOK, resp)
}

func CommentList(c *gin.Context) {
	var resp model.CommentListResponse
	//token := c.Query("token")
	video_string := c.Query("video_id")
	videoId, _ := strconv.Atoi(video_string)
	resp = handler.HandleCommentListGet(videoId)
	c.JSON(http.StatusOK, resp)
}
