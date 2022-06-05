package repository

import (
	"reflect"
	"simpledy/model"
)

var video = model.Video{}
var k = reflect.ValueOf(&video).Elem()
var varVideoId string
var varFavoriteCount string

//更改数据库状态
//根据ID找到对应视频,点赞数加/减1
func IfFavoriteActionYes(token string, video_id string, action_type string) (c bool) {
	c = false
	varVideoId = k.Type().Field(1).Name
	varFavoriteCount = k.Type().Field(5).Name
	var result1 int
	//根据视频ID找到其点赞总数
	db.Select(varFavoriteCount).Where(varVideoId+" = ?", video_id).Find(&result1)
	//预设结果
	result11 := result1 + 1
	result12 := result1 - 1
	if action_type == "1" {
		db.Model(&model.Video{}).Select(varFavoriteCount).Where(varVideoId+" = ?", video_id).Updates(result11)
		c = true
	} else {
		db.Model(&model.Video{}).Select(varFavoriteCount).Where(varVideoId+" = ?", video_id).Updates(result12)
		c = true
	}
	return c
}

//获取数据库信息
//找出User
func GetVideoList(user_id string, token string) {

}
