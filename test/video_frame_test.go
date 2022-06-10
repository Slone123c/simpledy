package test

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestScreenShot(t *testing.T) {
	filename := "H:\\simpledy-0.2.0\\static_data\\VIDEO_20220609_140404794.MP4"
	width := 640
	height := 360
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		panic("could not generate frame")
	}
	// Do something with buffer, which contains a JPEG image
}
