package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type tsts struct {
	ID   int
	Name string
	Abc  string
}

func main() {
	db := sqlConnect()
	db.AutoMigrate(&tsts{})
	defer db.Close()

	db.Create(&tsts{
		ID:   4,
		Name: "ddd",
		Abc:  "aiu",
	})
}

func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "root"
	PASS := "cts-renewal-db-password-BANANA"
	PROTOCOL := "tcp(mysql-db:3306)"
	DBNAME := "Test"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}

	fmt.Println("DB接続成功")

	return db
}
