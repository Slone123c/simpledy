package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simpledy/handler"
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

}
