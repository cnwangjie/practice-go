package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fin, _ := os.Open("./filtered_words.txt")
	buf := bufio.NewReader(fin)
	var words []string
	for {
		text, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		words = append(words, string(text))
	}
	var input string
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		}
		b := false
		for _, v := range words {
			if v == input {
				b = true
				break
			}
		}
		if b {
			fmt.Println("Freedom")
		} else {
			fmt.Println("Human Rights")
		}

	}
}
