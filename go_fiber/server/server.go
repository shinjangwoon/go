package main

import (
	"github.com/shinjangwoon/go/src/go_study/database"

	"github.com/gofiber/fiber/v2"
)

func init() {
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

	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"message": "Welcome to Myfirst Web Application"})

	})

	app.Listen(":8000")
}
