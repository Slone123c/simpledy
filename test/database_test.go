package test

import (
	"simpledy/driver"
	"testing"
)

//var user = model.User{
//	Username: "testUser",
//	Password: "123456",
//}

func TestDatabase(t *testing.T) {
	driver.InitData()
	//fmt.Println(db.RowsAffected)
}
