package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simpledy/handler"
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
