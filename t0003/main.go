package main

import (
	"bufio"
	"fmt"
	"os"

	"gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	fin, _ := os.Open("./codes.txt")
	buf := bufio.NewReader(fin)
	for {
		text, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		client.LPush("codes", text[:6])
	}
	fmt.Println("done!")

}
