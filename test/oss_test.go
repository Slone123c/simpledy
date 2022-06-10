package test

import (
	"fmt"
	"os"
	"path/filepath"
	"simpledy/utils"
	"testing"
)

func TestUploadFile(t *testing.T) {
	abs, _ := os.Getwd() // 获绝对路径
	fmt.Println(abs)
	dir := filepath.Dir(abs) // 上级目录
	fmt.Println("now:", dir)

	newdir := filepath.Join(dir, "static_data", "videos", "dog.MP4")
	fmt.Println("文件的路径为:", newdir)
	objectPath := "videos/TestingVideo.MP4" // bucket下 目录 +文件名
	localFilePath := newdir
	utils.UploadFile(localFilePath, objectPath)
}
