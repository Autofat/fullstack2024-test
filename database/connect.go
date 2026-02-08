package database

import (
	"PeepL-Test/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "host=localhost user=postgres password=2512 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Failed Connect")
	}
	err = DB.AutoMigrate(&models.My_client{})
	if err != nil{
		panic("Failed Migrate")
	}
	fmt.Println("Database Connected & migrated")
}