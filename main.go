package main

import (
	db "example.com/prac02/database"
	"example.com/prac02/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	db.ConnectDB()

	routes.SetupRoutes(app, db.DB)
	app.Listen(":8080")
}

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "OK"})
}
