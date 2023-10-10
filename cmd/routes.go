package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/greybluesea/dockerised-fullstack-webapp-gofiber-gorm-postgres/database"
	"github.com/greybluesea/dockerised-fullstack-webapp-gofiber-gorm-postgres/models"
)

func setupRoutes(app *fiber.App) {
	app.Get("/hello", helloHandler)
	app.Get("/", homeHandler)
	app.Get("/newfact", newfactHandler)
	app.Post("/create", createHandler)
	app.Post("/delete", deleteHandler)
}

func helloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World! Tony here 👋!!")
}
func homeHandler(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Find(&facts)
	// 	return c.Status(200).JSON(facts)
	return c.Render("index", fiber.Map{"Title": "Fun Facts", "Subtitle": "Dockerised Fullstack WebApp(GoFiber + GORM + Postgres) - learned from Div Rhino", "Facts": facts})
}

func newfactHandler(c *fiber.Ctx) error {
	return c.Render("newfact", fiber.Map{"Title": "New Fact", "Subtitle": "Add an interesting fact"})
}

func createHandler(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Create(&fact)
	//return c.Status(200).JSON(fact)
	return successHandler(c)
}

func successHandler(c *fiber.Ctx) error {
	return c.Render("success", fiber.Map{
		"Title": "Fact added successfully",
	})
}

func deleteHandler(c *fiber.Ctx) error {
	// Retrieve the fact ID from the POST request
	factID := c.FormValue("ID")
	result := database.DB.Delete(&models.Fact{}, factID)
	if result.Error != nil {
		return result.Error
	}
	// Redirect to the home page after deletion
	c.Redirect("/")
	return nil
}
