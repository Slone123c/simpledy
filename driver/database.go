package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simpledy/model"
)

// 初始化和数据的连接，不对数据做任何操作
func InitDB() (*gorm.DB, error) {
	fmt.Println("开始连接数据库")
	// 建立数据库连接
	host := "127.0.0.1"
	port := "3306"
	DBName := "simpledy"
	username := "root"
	password := "root"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		DBName,
		charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("连接数据库成功............")

	return db, err
}

func InitTable() {
	// 数据库自动建表
	db, _ := InitDB()
	//建立 User 表
	db.AutoMigrate(&model.User{})
	//建立 UserInformation 表
	db.AutoMigrate(&model.UserInformation{})
	//建立Video表
	db.AutoMigrate(&model.Video{})
	//建立Favorite表
	//db.AutoMigrate(&model.Favorite{})
	// m4 := db.Migrator()
	// err4 := m4.CreateTable(&model.Comment{})
	// if err4 != nil {
	// 	fmt.Println("建立Comment表失败")
	// } else {
	// 	fmt.Println("建立Comment表成功......")
	// }
	// db.Create(&comment)
}

// 用来注入一些初始数据来方便测试
func InitData() {
	db, _ := InitDB()
	var user = model.User{
		Id:       0,
		Username: "testUser",
		Password: "123456",
	}

	var userInfo = model.UserInformation{
		Id:            1,
		UserId:        0,
		Name:          "testInfo",
		FollowCount:   5,
		FollowerCount: 6,
	}

	db.Create(&user)
	db.Create(&userInfo)
}
