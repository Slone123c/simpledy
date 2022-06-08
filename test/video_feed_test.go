package test

import (
	"fmt"
	"simpledy/model"
	"simpledy/repository"
	"testing"
	"time"
)

var video = model.Video{
	AuthorId:      12,
	PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
	CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
	FavoriteCount: 0,
	CommentCount:  0,
	IsFavorite:    false,
	Title:         "TestingVideo",
	CreatedAt:     time.Now().Unix(),
}

func TestSth(t *testing.T) {
	for i := 0; i < 30; i++ {
		repository.InsertNewVideo(video)
	}

	//assert.Equal(t, 1, res)
}

func TestFindVideosBefore(t *testing.T) {
	resSet, res := repository.FindVideosBefore(1654667124)

	fmt.Println(resSet[0])
	fmt.Println(res)
}

//func TestHandleVideoFeedGet(t *testing.T) {
//	handler.HandleVideoFeedGet("1654667124", "")
//}
