package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
)

func main() {
	// 读原图片
	r, _ := os.Open("./i.png")
	img, _ := png.Decode(r)
	r.Close()

	// 读字体文件
	fontBytes, _ := ioutil.ReadFile("./arial.ttf")
	f, _ := freetype.ParseFont(fontBytes)

	// 建立一个新的空白图片
	m := image.NewRGBA(img.Bounds())
	draw.Draw(m, m.Bounds(), img, image.ZP, draw.Src)
	fg := image.NewUniform(color.RGBA{0xdd, 0x0, 0x0, 0xff})
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(48)
	c.SetClip(m.Bounds())
	c.SetDst(m)
	c.SetSrc(fg)
	pt := freetype.Pt(150, 50)

	// 写数字
	c.DrawString("4", pt)

	// 保存图片
	out, _ := os.Create("o.png")
	b := bufio.NewWriter(out)
	png.Encode(b, m)
	b.Flush()
	fmt.Println("done!")
}
