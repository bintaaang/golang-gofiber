package handlers

import (
    "gofiber-endpoint/database"
    "gofiber-endpoint/models"
    "gofiber-endpoint/middleware"

    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
)

// REGISTER Admin / Kurir
func Register(c *fiber.Ctx) error {
    type Request struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Phone    string `json:"phone"`
        Role     string `json:"role"`     // "admin" atau "courier"
        Password string `json:"password"` // plaintext, akan di-hash
    }

    var body Request
    if err := c.BodyParser(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    // Cek role valid
    if body.Role != "admin" && body.Role != "courier" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "role harus admin atau courier"})
    }

    // Hash password
    hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "gagal hash password"})
    }

    user := models.User{
        Name:  body.Name,
        Email: body.Email,
        Phone: body.Phone,
        Role:  body.Role,
        // simpan password hash di field baru di models User
        // kita tambahkan Password string di models User
        Password: string(hash),
    }

    database.UsingPostgre.Create(&user)

    return c.JSON(fiber.Map{
        "message": "user berhasil dibuat",
        "user_id": user.ID,
    })
}

// LOGIN Admin / Kurir
func Login(c *fiber.Ctx) error {
    type Request struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    var body Request
    if err := c.BodyParser(&body); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    var user models.User
    if err := database.UsingPostgre.First(&user, "email = ?", body.Email).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "email tidak ditemukan"})
    }

    // Cek password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "password salah"})
    }

    // Generate JWT
    token, err := middleware.GenerateToken(user.ID, user.Role)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "gagal generate token"})
    }

    return c.JSON(fiber.Map{
        "token": token,
    })
}
