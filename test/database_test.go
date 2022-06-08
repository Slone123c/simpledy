package test

import (
	"simpledy/driver"
	"simpledy/model"
	"testing"
)

//var user = model.User{
//	Username: "testUser",
//	Password: "123456",
//}

func TestDatabase(t *testing.T) {
	driver.InitTable()
	driver.InitData()
	//fmt.Println(db.RowsAffected)
}

func TestCustomDb(t *testing.T) {
	db, _ := driver.InitDB()
	db.AutoMigrate(&model.Video{})
}
