package controllers

import (
	"github.com/AhmadZ571/GOLANG.git/connection"
	"github.com/AhmadZ571/GOLANG.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	var user []models.User

	connection.DB.Find(&user)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All user",
		"data":    user,
	})

}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	connection.DB.Find(&user, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User found",
		"data":    user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	userInput := new(models.UserInput)

	if err := c.BodyParser(user); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	connection.DB.Create(&user)

	userInput.Name = user.Name
	userInput.Email = user.Email
	userInput.Password = user.Password

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User created",
		"data":    userInput,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user := new(models.User)

	userInput := new(models.UserInput)

	if err := c.BodyParser(userInput); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	connection.DB.Model(&user).Where("id = ?", id).Updates(userInput)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully updated",
		"data":    userInput,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	connection.DB.First(&user, id)

	if user.Name == "" {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "No user found with ID",
			"data":    nil,
		})
	}

	connection.DB.Delete(&user)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully deleted",
	})
}
