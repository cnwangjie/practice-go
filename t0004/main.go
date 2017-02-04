package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fin, _ := os.Open("./text.txt")
	buf := bufio.NewReader(fin)
	count := make(map[string]int)
	reg := regexp.MustCompile("[A-Za-z]+")
	for {
		text, err := buf.ReadString(' ')
		if err != nil {
			break
		}
		for _, v := range reg.FindAllString(text, -1) {
			v = strings.ToLower(v)
			count[v] += 1
		}
	}
	for text := range count {
		fmt.Println(text, count[text])
	}
}
