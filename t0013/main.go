package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
)

func main() {
	// 获取网页
	resp, _ := http.Get("http://tieba.baidu.com/p/2166231880")
	body, _ := ioutil.ReadAll(resp.Body)

	// 获取图片链接
	re, _ := regexp.Compile("http://imgsrc.baidu.com/forum/[!-~]*.jpg")
	urls := re.FindAllString(string(body), -1)
	total := len(urls)
	fmt.Println(total, "images in total")

	// 下载图片
	for n, v := range urls {
		tmpRes, _ := http.Get(v)
		dst, _ := os.Create("./saved/" + path.Base(v))
		io.Copy(dst, tmpRes.Body)
		fmt.Println("(", n+1, "/", total, ") done:", v)
	}
	fmt.Println("done!")
}
