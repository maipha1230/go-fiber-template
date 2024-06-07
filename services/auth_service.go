package services

import (
	"example.com/prac02/models"
	"example.com/prac02/repositories"
	"example.com/prac02/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (s *AuthService) Register(ctx *fiber.Ctx) error {
	body := models.SignUp{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, err.Error())
	}

	validate := utils.NewValidator()
	if err := validate.Struct(body); err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, utils.ValidatorErrors(err))
	}

	if existUser, err := s.userRepository.FindByEmail(body.Email); existUser != nil && err == nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, "Email already exist")
	}

	user := models.User{
		Email:    body.Email,
		Password: utils.GeneratePassword(body.Password),
	}
	if err := s.userRepository.Save(&user); err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"message":    "Sign Up Success",
	})
}

func (s *AuthService) SignIn(ctx *fiber.Ctx) error {
	body := models.SignIn{}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, err.Error())
	}

	validate := utils.NewValidator()
	if err := validate.Struct(body); err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, utils.ValidatorErrors(err))
	}

	existUser, err := s.userRepository.FindByEmail(string(body.Email))
	if existUser == nil && err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, "Email or Password not valid")
	}

	compareUserPassword := utils.ComparePasswords(existUser.Password, body.Password)
	if !compareUserPassword {
		return utils.ThrowExceoption(ctx, fiber.StatusBadRequest, "Email or Password not valid")
	}

	token, err := utils.GenerateJWT(existUser.ID, existUser.Email)
	if err != nil {
		return utils.ThrowExceoption(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"message":    "Sign In Success",
		"token": fiber.Map{
			"accessToken": token,
		},
	})

}
