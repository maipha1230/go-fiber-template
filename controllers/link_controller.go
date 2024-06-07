package controllers

import (
	"example.com/prac02/repositories"
	"example.com/prac02/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LinkController struct {
	linkService *services.LinkService
}

func NewLinkController(db *gorm.DB) *LinkController {
	linkRepository := repositories.NewLinktreeRepository(db)
	linkService := services.NewLinkService(linkRepository)
	return &LinkController{linkService: linkService}
}

func (controller *LinkController) CreateLink(ctx *fiber.Ctx) error {
	return controller.linkService.CreateLink(ctx)
}

func (controller *LinkController) UpdateLink(ctx *fiber.Ctx) error {
	return controller.linkService.UpdateLink(ctx)
}

func (controller *LinkController) GetLinksByUser(ctx *fiber.Ctx) error {
	return controller.linkService.GetLinksByUser(ctx)
}

func (controller *LinkController) DeleteLink(ctx *fiber.Ctx) error {
	return controller.linkService.DeleteLink(ctx)
}
