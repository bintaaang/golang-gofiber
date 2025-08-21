package middleware

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"fmt"
	"os"
	//"strings"
	"github.com/joho/godotenv"
)

func Load() []string {

	_ = godotenv.Load();
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

		//clientIP := ctx.Get("X-Forwarded-For")
		clientIP := "182.253.55.189:1082"
		fmt.Println(clientIP)
		if clientIP != "" {
			//clientIP = ctx.IP()
			if strings.Contains(clientIP, ":") {
				clientIP = strings.Split(clientIP, ":")[0]
			}
		} else {
			clientIP = strings.Split(clientIP, ",")[0]
			clientIP = strings.TrimSpace(clientIP)
			if strings.Contains(clientIP, ":") {
				clientIP = strings.Split(clientIP, ":")[0]
			}
		}
		fmt.Println("request dari ip :", clientIP)
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
