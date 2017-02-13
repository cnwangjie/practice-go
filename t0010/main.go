package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/golang/freetype"
)

func main() {
	fontBytes, _ := ioutil.ReadFile("./arial.ttf")
	f, _ := freetype.ParseFont(fontBytes)

	// 创建画布
	m := image.NewRGBA(image.Rect(0, 0, 100, 25))

	// 创建随机数生成器
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 设置freetype
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(25)
	c.SetClip(m.Bounds())
	c.SetDst(m)

	// 先在图片上写四个字母
	str := "qwertyuiopasdfghjklzxcvbnm"
	for n, _ := range [4]byte{} {
		c.SetSrc(image.NewUniform(color.RGBA{uint8(rander.Intn(255)), uint8(rander.Intn(255)), uint8(rander.Intn(255)), 0xff}))
		c.DrawString(string(str[rander.Intn(len(str))]), freetype.Pt(25*n, 20))
	}

	// 把空白用随机颜色填充
	for x := 0; x < m.Bounds().Max.X; x += 1 {
		for y := 0; y < m.Bounds().Max.Y; y += 1 {
			r, g, b, a := m.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 && a == 0 {
				m.Set(x, y, color.RGBA{uint8(rander.Intn(255)), uint8(rander.Intn(255)), uint8(rander.Intn(255)), 0xff})
			}
		}
	}

	// 这里最好加上高斯模糊，否则达不到题目的效果

	o, _ := os.Create("./result.png")
	png.Encode(o, m)
	o.Close()
	fmt.Println("done!")
}
