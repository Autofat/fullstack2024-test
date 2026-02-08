package main

import (
	"PeepL-Test/database"
	"PeepL-Test/pkg/redis"
	"PeepL-Test/routes"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()
	redis.RedisInit()

	app:=fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}