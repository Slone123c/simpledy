package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simpledy/driver"
	"simpledy/model"
)

func Register(c *gin.Context) {
	fmt.Println("Registering....")
	db, err := driver.InitDB()
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println("username is " + username)
	if err != nil {
		log.Fatal("Failed to connect to database, error:" + err.Error())
	}
	user := model.User{
		Name:     username,
		Password: password,
	}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功！",
	})

	//db.Where("username = ?", username).First(&users)
}
