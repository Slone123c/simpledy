package repository

import (
	"reflect"
	"simpledy/model"
)

/*
type UserInformation struct {
	Id            int64  `gorm:"size:64;not null"`
	Name          string `gorm:"size:32;not null"` // 用户名称
	FollowCount   int64  `gorm:"size:64;not null"`
	FollowerCount int64  `gorm:"size:64;not null"`
	// 使得 UserId 成为外键并关联至 User 表
	User   User `gorm:"ForeignKey:UserId"`
	UserId int64
}

/*
//var userInfo = model.UserInformation{}
//var db, _ = driver.InitDB()
// // 用于取得数据库对象名称
i = 0, Id
i = 1, Name
i = 2, FollowCount
i = 3, FollowerCount
i = 4, User
i = 5, UserId
varName = euser.Type().Field(i).Name
*/
var euserInfo = reflect.ValueOf(&model.UserInformation{}).Elem()

func FindUserInfoByUserId(userId int) model.UserInformation {
	var result model.UserInformation
	db.Model(model.UserInformation{}).Where("user_id = ?", userId).Scan(&result)
	return result
}
