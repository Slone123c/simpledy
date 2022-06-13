package repository

import (
	"fmt"
	"simpledy/model"
)

var video = model.Video{}
var videoInsert = model.Favorite{}

//判断Favorite表中是否存在此条点赞信息
func IfFavoriteActionYes(userId int64, video_id int64, action_type int) (model.Favorite, int) {
	var result model.Favorite
	res := db.Model(model.Favorite{}).Where("user_id = ? AND video_id = ?", userId, video_id).Scan(&result)
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

	res := db.Where("user_id = ? AND video_id = ?", userId, video_id).Unscoped().Delete(&videoInsert)
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
	var result1 model.Video
	//根据视频ID找到其点赞总数
	db.Model(&model.Video{}).Where("id = ?", video_id).Scan(&result1)

	//预设结果
	result12 := result1.FavoriteCount - 1
	fmt.Println(result12)

	if action_type == 1 {
		UpdateVideoFavoriteNumberPlusOneByVideoId(int(video_id))
		c = true
	} else if result12 >= 0 {
		UpdateVideoFavoriteNumberMinusOneByVideoId(int(video_id))
		c = true
	}
	return c
}

//获取用户喜欢的所有Video的Id
func FindFavoriteListIdByUserId(userId int64) ([]model.Favorite, int) {
	var favoriteList []model.Favorite
	res := db.Debug().Model(model.Favorite{}).Where("user_id = ?", userId).Find(&favoriteList)
	return favoriteList, int(res.RowsAffected)
}

//根据视频Id获取到视频对应的信息
func FindVideoInfoByVideoId(videoId int) (model.Video, int) {
	var video model.Video
	res := db.Debug().Model(model.Video{}).Where("id = ?", videoId).Find(&video)
	return video, int(res.RowsAffected)
}

//根据视频信息获取到对应的作者信息
func FindAuthorInfoByVideoInfo(allVideoInfo []model.Video) ([]model.UserInformation, int) {
	var authors []model.UserInformation
	res := db.Debug().Model(model.UserInformation{}).Find(&authors)
	return authors, int(res.RowsAffected)
}
