package repository

// import (
// 	"fmt"
// 	"reflect"
// 	"simpledy/model"
// )

// var relation = model.Relation{}
// var h = reflect.ValueOf(&relation).Elem()
// var varToUserId string

// //判断Relation表中是否存在此关注信息
// func IfRelationActionYes(user_id int64, touser_id int64) (model.Relation, int) {
// 	var result model.Relation
// 	varUserId = h.Type().Field(0).Name
// 	varToUserId = h.Type().Field(1).Name
// 	res := db.Model(model.Relation{}).Where(varUserId+" = ? AND"+varToUserId+" = ?", user_id, touser_id).Scan(&result)
// 	return result, int(res.RowsAffected)
// }

// //新生成Relation关注信息储存进到Relation表中
// func CreateRelationInformation(newRelationInformation model.Relation) int {
// 	res := db.Create(&newRelationInformation)
// 	fmt.Println("新关注信息已生成成功！")
// 	return int(res.RowsAffected)
// }

// //从Relation表中删除该条关注数据
// func DeleteRelationInformation(user_id int64, touser_id int64) bool {
// 	varUserId = h.Type().Field(0).Name
// 	varToUserId = h.Type().Field(1).Name
// 	res := db.Where(varUserId+" = ? AND"+varToUserId+" = ?", user_id, touser_id).Unscoped().Delete(&relation)
// 	affected := res.RowsAffected
// 	if affected > 0 {
// 		fmt.Println("数据成功删除")
// 		return true
// 	} else {
// 		fmt.Println("数据删除失败")
// 		return false
// 	}
// }

// //更改数据库状态
// //根据ID找到对应被关注/被取关用户,粉丝数量加/减1
// //根据ID找到对应操作用户,关注数量加/减1
// func UpdateRelationCount(user_id int64, touser_id int64, action_type string) (c bool) {
// 	c = false
// 	// var name = k.Type().Field(0).Name
// 	varFavoriteCount = k.Type().Field(5).Name
// 	var result1 int
// 	//根据ID找到用户的FollowerCount
// 	db.Select(varFavoriteCount).Where(varVideoId+" = ?", video_id).Find(&result1)
// 	//预设结果
// 	result11 := result1 + 1
// 	result12 := result1 - 1
// 	if action_type == "1" {
// 		db.Model(&model.Video{}).Select(varFavoriteCount).Where(varVideoId+" = ?", video_id).Updates(result11)
// 		c = true
// 	} else {
// 		db.Model(&model.Video{}).Select(varFavoriteCount).Where(varVideoId+" = ?", video_id).Updates(result12)
// 		c = true
// 	}
// 	return c
// }
