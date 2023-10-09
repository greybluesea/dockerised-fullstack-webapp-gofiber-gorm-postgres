package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/dockerised-rest-api-gofiber-postgres/database"
	"github.com/greybluesea/dockerised-rest-api-gofiber-postgres/models"
)

func setupRoutes(app *fiber.App) {
	app.Get("/hello", homeHandler)
	app.Post("/create", createHandler)
}

func homeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World! Tony here 👋!!")
}

func createHandler(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Create(&fact)

	return c.Status(200).JSON(fact)
}
