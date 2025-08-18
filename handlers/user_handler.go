package handlers

import (
	"gofiber-endpoint/database"
	"gofiber-endpoint/models"
	"gofiber-endpoint/models/request"

	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(ctx *fiber.Ctx) error {

	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {

		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func CreateDataUser(ctx *fiber.Ctx) error {

	user := new(request.UserCreateRequest) //menyiapkan

	if err := ctx.BodyParser(user); err != nil { //megubah format json ke format go

		return err
	}

	newUser := models.User{

		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
		Apikey: user.Apikey,
	}

	errCreateUser := database.DB.Create(&newUser).Error

	if errCreateUser != nil { //ini tetap berjalan karena tidak ada aturan di databasenya

		return ctx.Status(500).JSON(fiber.Map{

			"message": "failed to store data",
		})

	}
	return ctx.JSON(fiber.Map{

		"message": "sucess",
		"data":    newUser,
	})
}

func GetUserId(ctx *fiber.Ctx) error {

	userId := ctx.Params("id")
	var user models.User
	err := database.DB.First(&user, userId).Error
	if err != nil {

		return ctx.Status(404).JSON(fiber.Map{

			"message": "user not found",
		})
	}
	return ctx.JSON(fiber.Map{

		"message": "success",
		"data":    user,
	})

}

func UpdateUser(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	// 2. Cari user di database
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "User tidak ditemukan",
		})
	}

	// 3. Bind data dari request
	var input request.UserUpdateRequest
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	// 4. Update data user
	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Address != "" {
		user.Address = input.Address
	}
	if input.Phone != "" {
		user.Phone = input.Phone
	}

	// 5. Simpan ke database
	database.DB.Save(&user)

	// 6. Kirim response
	return ctx.JSON(fiber.Map{
		"message": "Update berhasil",
		"data":    user,
	})
}

func DeleteUserId(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	// Cek apakah user ada sebelum dihapus
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Soft delete
	err := database.DB.Delete(&user)
	if err != nil {
		return ctx.Status(200).JSON(fiber.Map{

			"message": "user telah di hapus",
		})
	}

	return ctx.SendStatus(204)
}
