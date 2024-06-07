package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ThrowExceoption(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"statusCode": status,
		"message":    message,
	})
}
