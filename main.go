package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kingztech2019/blogbackend/database"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files..")
	}

	database.Connect()

	port := os.Getenv("PORT")
	app := fiber.New()
	app.Listen(":" + port)

}
