package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randString(n int, len int) []string {
	// 根据当前时间产生一个随机数生成器
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]string, n)
	for i, _ := range result {
		// 生成一个随机数并求出它的md5值
		randSeed := md5.Sum([]byte(strconv.FormatInt(rander.Int63(), 10)))
		// 将md5值base64编码
		text := base64.StdEncoding.EncodeToString(randSeed[0:])
		fmt.Println(text)
		result[i] = text[0:len]
	}
	return result
}

func main() {
	codes := randString(200, 6)

	fout, _ := os.Create("./codes.txt")
	for _, v := range codes {
		fmt.Println(v)
		fout.WriteString(v + "\n")
	}
	fout.Close()
}
