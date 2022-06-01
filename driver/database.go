package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simpledy/model"
)

func InitDB() (*gorm.DB, error) {
	fmt.Println("开始连接数据库")
	host := "localhost"
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

	db.AutoMigrate(&model.User{}) // 数据库自动建表

	return db, err
}
