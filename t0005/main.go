package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/nfnt/resize"
)

func changeResolution(name string) {
	r, _ := os.Open("./images/" + name)
	img, _ := png.Decode(r)
	tmp, _ := png.DecodeConfig(r)
	r.Close()
	height := float64(tmp.Height)
	width := float64(tmp.Width)

	// 计算等比缩放后的大小
	if height > width {
		if height/width > 1.775 {
			width *= 1136 / height
			height = 1136
		} else {
			height *= 640 / width
			width = 640
		}
	} else {
		if width/height > 1.775 {
			height *= 1136 / width
			width = 1136
		} else {
			width *= 640 / height
			height = 640
		}
	}

	// resize
	cvs := resize.Thumbnail(uint(width), uint(height), img, resize.Lanczos3)

	o, _ := os.Create("./images/changed" + name)
	png.Encode(o, cvs)
	o.Close()
}

func main() {
	// 读取目录里的文件
	images, _ := ioutil.ReadDir("./images")
	for _, v := range images {
		changeResolution(v.Name())
	}
}
