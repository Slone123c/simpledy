package repository

import (
	"fmt"
	"simpledy/model"
)

func InsertNewVideo(video model.Video) int {
	res := db.Create(&video)
	fmt.Println("新视频创建成功！", "插入新数据", res.RowsAffected, "条")
	return int(res.RowsAffected)
}

func FindVideosBefore(time int) ([]model.Video, int) {
	var videos []model.Video
	res := db.Debug().Limit(30).Model(model.Video{}).Order("created_at DESC").Where("created_at <= ?", time).Find(&videos)
	return videos, int(res.RowsAffected)
}
