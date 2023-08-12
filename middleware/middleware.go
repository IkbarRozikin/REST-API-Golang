package middleware

import (
	"go-fiber-gorm/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Unauthorized",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Unauthorized",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		// userInfo := ctx.Locals("userInfo")
		// log.Println("user info data:", userInfo)
		log.Println("info data:", claims)
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"Message": "Forbiden Access",
		})

	}
	log.Println("info data:", claims)

	return ctx.Next()
}
