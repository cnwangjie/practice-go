package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test?charset=utf8")

	db.Exec("CREATE TABLE `codes`(`code` VARCHAR(255) NOT NULL);")

	fin, _ := os.Open("./codes.txt")
	buf := bufio.NewReader(fin)
	tx, _ := db.Begin()
	for {
		text, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		tx.Exec("INSERT INTO codes(code) VALUES(?)", text[:6])
	}
	tx.Commit()
	fmt.Println("done!")
	db.Close()
}
