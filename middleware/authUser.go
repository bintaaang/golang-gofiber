package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
    token := c.Get("Authorization") // misal Bearer <token>
    if token == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
    }

    // decode token, dapatkan userID dan role
    userID, role, err := ParseJWT(token)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
    }

    // simpan ke locals untuk controller
    c.Locals("userID", userID)
    c.Locals("role", role)

    return c.Next()
}
