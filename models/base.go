package models

import (
	"fmt"
	"github.com/ichtrojan/thoth"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
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


	dbDetails := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",username,password, dbHost, dbName) //Build connection string
	

	conn, err := gorm.Open("mysql", dbDetails)
	fmt.Println(dbDetails)

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}

	db = conn

	db.DropTable(&Events{}, &EventKind{})

	db.Debug().AutoMigrate(&Users{}, &Events{}, &EventKind{})

	events := []Events{
		{
			Name:    "Wizkid (Made in Lagos)",
			Address: "Eko Hotel",
			Code:    "wc-20",
			Image:   "https://res.cloudinary.com/siteat/image/upload/v1590522733/b04760a8edb1b133ac28bcf922a1b29d_qmiiow.jpg",
			Details: "Wizkid concert ...",
			Date:    time.Now(),
			Event_Kinds: []EventKind{
				{
					Name : "VIP",
					Price: 2000,
				},
				{
					Name:  "VVIP",
					Price: 3000,
				},
			},
		},
		{
			Name:    "Davido (Made in Accra)",
			Address: "Eko Atlantic",
			Code:    "da-10",
			Image:   "https://res.cloudinary.com/siteat/image/upload/v1590522942/Davido_bf6ceh.jpg",
			Details: "Davido concert OBO...",
			Date  :  time.Now(),
			Event_Kinds: []EventKind{
				{
					Name : "Table for 4",
					Price: 10000,
				},
				{
					Name:  "Regular",
					Price: 1000,
				},
			},
		},
	}

	for _, event := range events {
		db.Create(&event)
	}

}

func GetDB() *gorm.DB {
	return db
}