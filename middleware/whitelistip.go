package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"

	//"strings"
	"github.com/joho/godotenv"
)

func Load() []string {

	_ = godotenv.Load()
	// if err := godotenv.Load(); err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"code": "500",
	// 		"message": "error",
	// 	})
	// }
	whitelist := os.Getenv("WHITELIST_IPS")
	// whitelist := []string{
	// 	"127.0.0.1",
	// 	"171.201.20.1",
	// }
	if whitelist == "" {
		return []string{}
	}
	// split by koma
	return strings.Split(whitelist, ",")
}

func NewIPWhitelistMiddleware(whitelist []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		clientIP := ctx.Get("X-Forwarded-For")
		if clientIP != "" {
			clientIP = strings.Split(clientIP, ",")[0]
		} else {
			clientIP = ctx.IP()
		}
		clientIP = strings.Split(clientIP, ":")[0] // hapus port kalau ada
		clientIP = strings.TrimSpace(clientIP)

		fmt.Println("Request dari IP:", clientIP)

		allowed := false
		for _, ip := range whitelist {
			if clientIP == ip {
				allowed = true
				break
			}
		}

		if !allowed {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Access denied",
				"ip":    clientIP,
			})
		}
		return ctx.Next()
	}
}
