package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/shinjangwoon/go/go_fiber/database"
	"github.com/shinjangwoon/go/go_fiber/router"
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file")
	}

	database.ConnectDB()
}

func main() {

	database.ConnectDB()

	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection: ")
	}

	defer sqlDb.Close()

	app := fiber.New()

	// Or extend your config for customization

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":8000")
}
