package handler

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGet(ctx *fiber.Ctx) error {

	var user []entity.User

	result := database.DB.Find(&user)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(user)
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id=?", userId).Error

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Message": "id not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"Message": "Succsess",
		"data":    user,
	})
}

func UserHandlerCreat(ctx *fiber.Ctx) error {
	user := new(request.CreatUserRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()

	errValidate := validate.Struct(*user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Message": "Failed",
			"error":   errValidate.Error(),
		})

	}

	newUser := entity.User{
		Nama:   user.Nama,
		Email:  user.Email,
		Adress: user.Address,
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Messaage": "Failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"Message": "Succes",
		"data":    newUser,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {

	userRequest := new(request.UpdateUserRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Message": "Bad Request",
		})
	}

	var user entity.User

	userId := ctx.Params("id")
	err := database.DB.First(&user, "id=?", userId).Error

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Message": "Kode not found",
		})
	}

	if userRequest.Nama != "" {
		user.Nama = userRequest.Nama
	}
	user.Email = userRequest.Email

	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Message": "Internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"Message": "Succsess",
		"data":    user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userKode := ctx.Params("id")
	var user entity.User

	err := database.DB.Debug().First(&user, "id=?", userKode).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"Message": "User not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error

	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Message": "Internal serve error",
		})
	}

	return ctx.JSON(fiber.Map{
		"Message": "User was deleted",
	})
}
