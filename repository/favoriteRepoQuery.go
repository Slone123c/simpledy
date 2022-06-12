package repository

import (
	"fmt"
	"reflect"
	"simpledy/model"
)

var video = model.Video{}
var videoInsert = model.Favorite{}
var k = reflect.ValueOf(&video).Elem()
var g = reflect.ValueOf(&videoInsert).Elem()
var varVideoId string
var varFavoriteCount string
var varUserId string

//判断Favorite表中是否存在此条点赞信息
func IfFavoriteActionYes(userId int64, video_id int64, action_type int) (model.Favorite, int) {
	var result model.Favorite
	// varUserId = g.Type().Field(0).Name
	// varVideoId = g.Type().Field(1).Name
	res := db.Model(model.Favorite{}).Where("user_id2 = ? AND video_id = ?", userId, video_id).Scan(&result)
	return result, int(res.RowsAffected)
}

//新生成Favorite点赞信息储存进到Favorite表中
func CreateFavoriteInformation(newFavoriteInformation model.Favorite) int {
	res := db.Create(&newFavoriteInformation)
	fmt.Println("新点赞信息已生成成功！")
	return int(res.RowsAffected)
}

//从Favorite表中删除该条点赞数据
func DeleteFavoriteInformation(userId int64, video_id int64) bool {
	varUserId = g.Type().Field(0).Name
	varVideoId = g.Type().Field(1).Name
	res := db.Where("user_id2 = ? AND video_id = ?", userId, video_id).Unscoped().Delete(&videoInsert)
	affected := res.RowsAffected
	if affected > 0 {
		fmt.Println("数据成功删除")
		return true
	} else {
		fmt.Println("数据删除失败")
		return false
	}
}

//更改数据库状态
//根据ID找到对应视频,点赞数加/减1
func UpdateFavoriteCount(token string, video_id int64, action_type int) (c bool) {
	c = false
	// varVideoId = k.Type().Field(0).Name
	// varFavoriteCount = k.Type().Field(5).Name
	var result1 model.Video
	//根据视频ID找到其点赞总数
	db.Model(&model.Video{}).Where("id = ?", video_id).Scan(&result1)
	//预设结果

	result11 := result1.FavoriteCount + 1
	fmt.Println(result11)
	// result11Str := strconv.Itoa(result11)
	result12 := result1.FavoriteCount - 1
	fmt.Println(result12)
	// result12Str := strconv.Itoa(result12)
	if action_type == 1 {
		db.Model(&model.Video{}).Where("id = ?", video_id).Update("favorite_count", result11)
		c = true
	} else {
		db.Model(&model.Video{}).Where("id = ?", video_id).Update("favorite_count", result12)
		c = true
	}
	return c
}

//获取用户喜欢的所有Video的Id
func FindVideoIdByUserId(user_id int64) ([]model.Favorite, int) {
	var record []model.Favorite
	res := db.Debug().Model(model.Favorite{}).Where("user_id2 = ?", user_id).Find(&record)
	return record, int(res.RowsAffected)
}

//根据视频Id获取到视频对应的信息
func FindVideoInfoByVideoId(allFavoriteRecord []model.Favorite) ([]model.Video, int) {
	var videos []model.Video
	res := db.Debug().Model(model.Video{}).Find(&videos)
	return videos, int(res.RowsAffected)
}

//根据视频信息获取到对应的作者信息
func FindAuthorInfoByVideoInfo(allVideoInfo []model.Video) ([]model.UserInformation, int) {
	var authors []model.UserInformation
	res := db.Debug().Model(model.UserInformation{}).Find(&authors)
	return authors, int(res.RowsAffected)
}
