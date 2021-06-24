package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"./assets"
)

func main() {
	db := sqlConnect()
	db.AutoMigrate(&assets.Tsts{})
	defer db.Close()

	db.Create(&assets.Tsts{
		ID:   4,
		Name: "ddd",
		Abc:  "aiu",
	})
}

func sqlConnect() (database *gorm.DB) {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"

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
