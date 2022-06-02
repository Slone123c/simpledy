package repository

import (
	"fmt"
	"reflect"
	"simpledy/driver"
	"simpledy/model"
)

/*
type User struct {
	gorm.Model
	Username string `gorm:"size:32;not null"`
	Password string `gorm:"size:32;not null"`
}

*/
var user = model.User{}
var db, _ = driver.InitDB()
var e = reflect.ValueOf(&user).Elem() // 用于取得数据库对象名称
/*
i = 0, gorm.Model
i = 1, Username
i = 2, Password
varName = e.Type().Field(i).Name
*/
var varName string

// 查询语句
func IfUserExistsById(id int) bool {
	result := db.First(&user, id)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func IfUserExistsByUsername(name string) bool {
	varName = e.Type().Field(1).Name
	result := db.Where(varName+" = ?", name).First(&user)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func FindUserByUserName(username string) model.User {
	var result model.User
	varName = e.Type().Field(1).Name
	db.Model(model.User{}).Where(varName+" = ?", username).Scan(&result)
	return result
}

// 增加语句
func InsertNewUser(user model.User) int {
	db.Create(&user)
	fmt.Println("新用户创建成功！", "用户名: "+user.Username)
	return int(db.RowsAffected)
}

// 删除语句
func DeleteUserByUsername(username string) bool {
	varName = e.Type().Field(1).Name
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
