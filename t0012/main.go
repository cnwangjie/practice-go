package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fin, _ := os.Open("./filtered_words.txt")
	buf := bufio.NewReader(fin)
	re := "["
	for {
		text, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		re += "(" + string(text) + ")"
	}
	re += "]"
	words, _ := regexp.Compile(re)
	var input string
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		}
		fmt.Println(words.ReplaceAllString(input, "*"))
	}
}
