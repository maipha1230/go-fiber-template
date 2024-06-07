package controllers

import (
	"example.com/prac02/repositories"
	"example.com/prac02/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(db *gorm.DB) *AuthController {
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	return &AuthController{authService: authService}
}

func (controller *AuthController) Register(ctx *fiber.Ctx) error {
	return controller.authService.Register(ctx)
}

func (controller *AuthController) SignIn(ctx *fiber.Ctx) error {
	return controller.authService.SignIn(ctx)
}
