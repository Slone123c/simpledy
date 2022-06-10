package utils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"strings"
)

func GenerateCoverSnapshot(videoFilePath, snapshotFilePath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoFilePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成封面图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成封面图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotFilePath+".png")
	if err != nil {
		log.Fatal("生成封面图失败：", err)
		return "", err
	}

	names := strings.Split(snapshotFilePath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}
