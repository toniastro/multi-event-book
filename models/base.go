package models

import (
	"fmt"
	"github.com/ichtrojan/thoth"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB 

func init() {
	logger, _ := thoth.Init("log")

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")


	db_details := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",username,password, dbHost, dbName) //Build connection string
	

	conn, err := gorm.Open("mysql", db_details)
	// fmt.Println(db_details)

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Users{}, &Events{}, &EventKind{})
}

func GetDB() *gorm.DB {
	return db
}