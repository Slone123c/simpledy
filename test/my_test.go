package test

import (
	"fmt"
	"simpledy/utils"
	"testing"
)

func TestTime(t *testing.T) {

	//obj := "https://simpledy.oss-cn-hangzhou.aliyuncs.com"
	//obj += "/videos"
	//obj += "/VIDEO_20220609_140404794.MP4"
	fmt.Println(utils.GetLocalPath())
	fmt.Println(utils.GetCoverPath())
	fmt.Println(utils.GetVideoPath())

	//https://simpledy.oss-cn-hangzhou.aliyuncs.com/videos/VIDEO_20220609_140404794.MP4
}
