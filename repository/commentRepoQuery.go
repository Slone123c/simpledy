package repository

import (
	"fmt"
	"simpledy/model"
)

// 增
func InsertNewComment(comment model.Comment) int {
	res := db.Create(&comment)
	fmt.Println("新评论创建成功！", "插入新数据", res.RowsAffected, "条")
	return int(res.RowsAffected)
}

// 查
func FindAllCommentsByVideoId(videoId int) ([]model.Comment, int) {
	var comments []model.Comment
	res := db.Debug().Model(model.Comment{}).Where("video_id = ?", videoId).Find(&comments)
	return comments, int(res.RowsAffected)
}

// 删
func DeleteCommentByVideoIdAndCommentID(videoId int, commentId int) int {
	res := db.Debug().Model(model.Comment{}).Where("video_id = ? and id = ?", videoId, commentId).Delete(&model.Comment{})
	return int(res.RowsAffected)
}
