package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/dockerised-rest-api-gofiber-postgres/database"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
