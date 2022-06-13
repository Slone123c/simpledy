package repository

import (
	"fmt"
	"gorm.io/gorm"
	"simpledy/model"
)

// 增加
func InsertNewVideo(video model.Video) int {
	res := db.Create(&video)
	fmt.Println("新视频创建成功！", "插入新数据", res.RowsAffected, "条")
	return int(res.RowsAffected)
}

// 查找
func FindLatestVideo() (model.Video, int) {
	var video model.Video
	res := db.Debug().Model(model.Video{}).Order("created_at DESC").First(&video)
	return video, int(res.RowsAffected)
}

func FindVideosBefore(time int) ([]model.Video, int) {
	var videos []model.Video
	res := db.Debug().Limit(30).Model(model.Video{}).Order("created_at DESC").Where("created_at <= ?", time).Find(&videos)
	return videos, int(res.RowsAffected)
}

func FindVideosAfter(time int) ([]model.Video, int) {
	var videos []model.Video
	res := db.Debug().Limit(30).Model(model.Video{}).Order("created_at DESC").Where("created_at >= ?", time).Find(&videos)
	return videos, int(res.RowsAffected)
}

func FindVideosByUserId(userId int) ([]model.Video, int) {
	var videos []model.Video
	res := db.Debug().Model(model.Video{}).Where("author_id = ?", userId).Find(&videos)
	return videos, int(res.RowsAffected)
}

// 改
func UpdateVideoCommentNumberPlusOneByVideoId(videoId int) {
	fmt.Println("评论数量更新！")
	db.Debug().Model(&model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", 1))
}
func UpdateVideoCommentNumberMinusOneByVideoId(videoId int) {
	fmt.Println("评论数量更新！")
	db.Debug().Model(&model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1))
}

func UpdateVideoFavoriteNumberMinusOneByVideoId(videoId int) {
	fmt.Println("评论数量更新！")
	db.Debug().Model(&model.Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
}
func UpdateVideoFavoriteNumberPlusOneByVideoId(videoId int) {
	fmt.Println("评论数量更新！")
	db.Debug().Model(&model.Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
}
