package handler

import (
	"fmt"
	"path/filepath"
	"simpledy/model"
	"simpledy/repository"
	"simpledy/utils"
	"time"
)

var snapshotPath = utils.GetCoverPath()
var videoPath = utils.GetVideoPath()
var url = "simpledy.oss-cn-hangzhou.aliyuncs.com"

func HandlePublishPost(videoFileName string, token string, title string) model.VideoPublishResponse {
	fmt.Println("封面路径:", snapshotPath)
	fmt.Println("视频路径", videoPath)

	userId := utils.ParseTokenAndGetUserId(token)
	var video model.Video
	// 取得本地视频和封面文件绝对路径
	videoFilePath := filepath.Join(videoPath, videoFileName) // 视频文件的完整目录
	fileName := utils.GetFilePrefix(videoFileName)           // 不包含后缀的视频文件名
	newCoverPath := filepath.Join(snapshotPath, fileName)    // 生成视频封面截图后的存放路径
	// 生成视频封面截图并取得封面图片文件名
	snapshotName, _ := utils.GenerateCoverSnapshot(videoFilePath, newCoverPath, 1)
	// 生成在 OSS数据库中的路径
	// OSS数据库中的文件夹路径
	videoObjectPath := "videos/"
	coverObjectPath := "covers/"
	// OSS数据库中的文件路径
	videoObjectPath = videoObjectPath + videoFileName // 例: videos/VIDEO_20220609_155429105.MP4
	coverObjectPath = coverObjectPath + snapshotName  // 例: covers/VIDEO_20220609_155429105.png
	//本地图片文件路径
	// 例如 H:\simpledy-0.2.0\static_data\covers\VIDEO_20220609_155429105.png
	snapshotFilePath := filepath.Join(snapshotPath, snapshotName)
	// 向 OSS数据库上传文件
	utils.UploadFile(snapshotFilePath, coverObjectPath)
	utils.UploadFile(videoFilePath, videoObjectPath)
	playUrl := url + "videos/" + videoFileName
	coverUrl := url + "covers/" + snapshotName
	fmt.Println("playUrl:", playUrl)
	fmt.Println("coverUrl:", coverUrl)
	video.CreatedAt = time.Now().Unix()
	video.AuthorId = int64(userId)
	video.Title = title
	video.PlayUrl = playUrl
	video.CoverUrl = coverUrl
	repository.InsertNewVideo(video)
	resp := model.GenerateVideoPublishResponse(0, "视频上传成功")
	return resp
}
