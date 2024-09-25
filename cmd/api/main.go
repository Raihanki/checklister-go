package main

import (
	"log"
	"os"

	"github.com/Raihanki/checklisters/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error when loading environtment file ERR::%v", err)
	}

	db := database.GetDatabaseConnection()

	app := fiber.New()
	Routes(app, db)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
