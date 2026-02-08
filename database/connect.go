package database

import (
	"PeepL-Test/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {

	dsn := "host=localhost user=postgres password=2512 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Failed Connect")
	}
	err = db.AutoMigrate(&models.My_client{})
	if err != nil{
		panic("Failed Migrate")
	}
	fmt.Println("Database Connected & migrated")
}