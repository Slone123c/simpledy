package utils

import (
	"os"
	"path"
	"path/filepath"
)

// 取得项目目录
func GetLocalPath() string {
	absPath, _ := os.Getwd()
	//dir := filepath.Dir(absPath)
	return absPath
}

func GetCoverPath() string {
	path := GetLocalPath()
	newPath := filepath.Join(path, "static_data", "covers")
	return newPath
}

func GetVideoPath() string {
	path := GetLocalPath()
	newPath := filepath.Join(path, "static_data", "videos")
	return newPath
}

func GetFilePrefix(filename string) string {
	filenameall := path.Base(filename)
	filesuffix := path.Ext(filename)
	fileprefix := filenameall[0 : len(filenameall)-len(filesuffix)]
	return fileprefix
}
