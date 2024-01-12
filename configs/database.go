package configs

import (
	"first-jwt/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/remember_it?parseTime=true"), &gorm.Config{})
	if err != nil{
		panic("Failed To Connect To Database")
	}

	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Database Connected")
}