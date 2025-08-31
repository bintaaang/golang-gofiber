package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
    token := ctx.Get("Authorization") // misal Bearer <token>
    if token == "" {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
    }

    // decode token, dapatkan userID dan role
    userID, role, err := ParseJWT(token)
    if err != nil {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
    }

    // simpan ke locals untuk controller
    ctx.Locals("userID", userID)
    ctx.Locals("role", role)

    return ctx.Next()
}
