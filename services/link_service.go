package services

import (
	model "example.com/prac02/models"
	"example.com/prac02/repositories"
	"example.com/prac02/utils"
	"github.com/gofiber/fiber/v2"
)

type LinkService struct {
	linkRepository repositories.LinktreeRepository
}

func NewLinkService(linkRepository repositories.LinktreeRepository) *LinkService {
	return &LinkService{linkRepository: linkRepository}
}

func (s *LinkService) CreateLink(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	body := model.LinkBodyRequest{}
	if err := ctx.BodyParser(&body); err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, err.Error())
	}

	validate := utils.NewValidator()
	if err := validate.Struct(body); err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, utils.ValidatorErrors(err))
	}

	link := model.Link{
		Title:  body.Title,
		Url:    body.Url,
		Type:   body.Type,
		UserID: userID,
	}
	err := s.linkRepository.CreateLink(&link)
	if err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"msg":        "Create Link Success",
	})

}

func (s *LinkService) UpdateLink(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)

	linkId, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, err.Error())
	}

	body := model.LinkBodyRequest{}
	if err := ctx.BodyParser(&body); err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, err.Error())
	}

	validate := utils.NewValidator()
	if err := validate.Struct(body); err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, utils.ValidatorErrors(err))
	}

	existLink, err := s.linkRepository.FindLinkByID(uint(linkId))
	if err != nil && existLink == nil {
		return utils.ThrowExceoption(ctx, fiber.StatusNotFound, "Link Not Found")
	}

	link := model.Link{
		Title:  body.Title,
		Url:    body.Url,
		Type:   body.Type,
		UserID: userID,
	}
	link.ID = uint(linkId)
	err = s.linkRepository.UpdateLink(&link)
	if err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"msg":        "Update Link Success",
	})
}

func (s *LinkService) GetLinksByUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)

	result, err := s.linkRepository.GetLinksByUser(userID)
	if err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusNotFound, "Links Not Found")
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"msg":        "Success",
		"data":       result,
	})
}

func (s *LinkService) DeleteLink(ctx *fiber.Ctx) error {
	linkId, err := ctx.ParamsInt("id")
	if err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, err.Error())
	}

	existLink, err := s.linkRepository.FindLinkByID(uint(linkId))
	if err != nil && existLink == nil {
		return utils.ThrowExceoption(ctx, fiber.StatusNotFound, "Links Not Found")
	}

	err = s.linkRepository.DeleteLink(existLink)
	if err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"msg":        "Delete Link Success",
	})

}
