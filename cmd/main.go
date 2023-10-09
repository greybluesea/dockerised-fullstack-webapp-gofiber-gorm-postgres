package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/dockerised-rest-api-gofiber-postgres/database"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! Tony here 👋!!!")
	})

	app.Listen(":3000")
}
