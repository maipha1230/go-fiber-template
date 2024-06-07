package routes

import (
	"example.com/prac02/controllers"
	"example.com/prac02/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	authController := controllers.NewAuthController(db)
	authGroup := app.Group("/auth")
	authGroup.Post("/signup", authController.Register)
	authGroup.Post("/signin", authController.SignIn)

	linkController := controllers.NewLinkController(db)
	linkGroup := app.Group("/link")
	linkGroup.Use(utils.JWTMiddleware)
	linkGroup.Post("/create", linkController.CreateLink)
	linkGroup.Post("/update/:id", linkController.UpdateLink)
	linkGroup.Get("/links", linkController.GetLinksByUser)
	linkGroup.Delete("/delete/:id", linkController.DeleteLink)
}
