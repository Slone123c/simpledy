package repository

import (
	"fmt"
	"reflect"
	"simpledy/driver"
	"simpledy/model"
)

/*
type User struct {
	Id       int64  `gorm:"size:64;not null"`
	Username string `gorm:"size:32;not null"`
	Password string `gorm:"size:32;not null"`
}


*/
var user = model.User{}
var db, _ = driver.InitDB()
var euser = reflect.ValueOf(&user).Elem() // 用于取得数据库对象名称
/*
i = 0, Id
i = 1, Username
i = 2, Password
varName = euser.Type().Field(i).Name
*/
var varName string

// 查询语句

func FindUserByUserName(username string) (model.User, int) {
	var result model.User
	varName = euser.Type().Field(1).Name
	res := db.Model(model.User{}).Where(varName+" = ?", username).Scan(&result)
	return result, int(res.RowsAffected)
}

func FindUserByUserId(userId int) (model.User, int) {
	var result model.User
	varName = euser.Type().Field(0).Name
	res := db.Model(model.User{}).Where(varName+" = ?", userId).Scan(&result)
	return result, int(res.RowsAffected)
}

// 增加语句
func InsertNewUser(user model.User) int {
	res := db.Create(&user)
	fmt.Println("新用户创建成功！", "用户名: "+user.Username)
	return int(res.RowsAffected)
}

func InsertNewUserInformation(userinfo model.UserInformation) int {
	res := db.Create(&userinfo)
	fmt.Println("新用户信息已创建！")
	return int(res.RowsAffected)
}

// 删除语句
func DeleteUserByUsername(username string) bool {
	varName = euser.Type().Field(1).Name
	res := db.Where(varName+" = ?", username).Unscoped().Delete(&user)
	affected := res.RowsAffected
	if affected > 0 {
		fmt.Printf("数据删除成功,删除了%v条数据。", affected)
		return true
	} else {
		fmt.Println("没有找到username为%v的数据,没有数据被删除。", username)
		return false
	}
}
