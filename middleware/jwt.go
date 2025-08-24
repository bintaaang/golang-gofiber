package middleware

import (
    "time"
    "errors"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
    "fmt"
    "strings"
)

var jwtSecret = []byte("MySuperSecretKey1234567890")

type JwtClaims struct {
    UserID uint   `json:"user_id"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

func GenerateToken(userID uint, role string) (string, error) {
    claims := JwtClaims{
        UserID: userID,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func Protected() fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        authHeader := ctx.Get("Authorization")
        if authHeader == "" {
            return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
            
        }
         parts := strings.Split(strings.TrimSpace(authHeader), " ")
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid auth header"})
        }
        tokenStr := parts[1]

        //token := authHeader; fmt.Println("Token string:", tokenStr)
        token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        if err != nil || !token.Valid {
            fmt.Println("Error parsing:", err)
            return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
        }

        claims, ok := token.Claims.(*JwtClaims)
        if !ok {
            return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
        }

        ctx.Locals("userID", claims.UserID)
        ctx.Locals("role", claims.Role)

        return ctx.Next()
    }
}

func RequireRole(role string) fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        userRole := ctx.Locals("role")
        if userRole != role {
            return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "access forbidden"})
        }
        return ctx.Next()
    }
}

func ParseJWT(tokenStr string) (uint, string, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil || !token.Valid {
        return 0, "", errors.New("invalid token")
    }

    claims, ok := token.Claims.(*JwtClaims)
    if !ok {
        return 0, "", errors.New("invalid claims")
    }

    return claims.UserID, claims.Role, nil
}
