package middlewares

import (
	"strings"
	"time"

	"github.com/CodeWithTamim/TaskManagementApi/dto"
	"github.com/CodeWithTamim/TaskManagementApi/repository"
	"github.com/CodeWithTamim/TaskManagementApi/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthHandler(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	authParsed := strings.Fields(authHeader)

	if len(authParsed) < 2 || strings.ToLower(authParsed[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
			Data:       nil,
			Success:    false,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized !",
		})
	}

	jwtToken := authParsed[1]

	email, exp, err := utils.ParseJWTClaim(jwtToken)

	if err != nil || time.Now().Unix() > exp {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
			Data:       nil,
			Success:    false,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized !",
		})
	}

	if !repository.UserExists(email) {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
			Data:       nil,
			Success:    false,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized !",
		})
	}

	c.Locals("email", email)

	return c.Next()
}
