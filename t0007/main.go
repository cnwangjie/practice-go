package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 直接读当前目录了 连着编译后的二进制文件一起读 很应垂死听
	files, _ := ioutil.ReadDir("./")
	line, note, empty := 0, 0, 0
	for _, v := range files {
		if v.IsDir() {
			continue
		}
		fin, _ := os.Open("./" + v.Name())
		buf := bufio.NewReader(fin)
		for {
			text, _, err := buf.ReadLine()
			if err != nil {
				break
			}
			// 去除空格和tab
			text = []byte(strings.Trim(string(text), " \t"))

			// 判断类型
			if len(text) == 0 {
				empty += 1
			} else if text[0] == '/' || text[0] == '*' {
				note += 1
			} else {
				line += 1
			}
		}
		fin.Close()
	}
	fmt.Println("Total", line, "line of code ", note, "line of note ", empty, "empty line")
}
