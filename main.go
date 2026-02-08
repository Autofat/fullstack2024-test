package main

import (
	"PeepL-Test/database"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.Connect()

	app:=fiber.New()
	log.Fatal(app.Listen(":3000"))
}