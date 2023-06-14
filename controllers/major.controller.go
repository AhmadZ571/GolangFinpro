package controllers

import (
	"github.com/AhmadZ571/GOLANG.git/connection"
	"github.com/AhmadZ571/GOLANG.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllMajor(c *fiber.Ctx) error {
	var major []models.Majors

	connection.DB.Find(&major)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All major",
		"data":    major,
	})
}

func GetMajor(c *fiber.Ctx) error {
	id := c.Params("id")
	var major models.Majors

	connection.DB.Find(&major, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Major found",
		"data":    major,
	})
}

func CreateMajor(c *fiber.Ctx) error {
	major := new(models.Majors)

	majorInput := new(models.MajorInput)

	if err := c.BodyParser(major); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	connection.DB.Create(&major)

	majorInput.Name = major.Name

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Major created",
		"data":    majorInput,
	})
}

func UpdateMajor(c *fiber.Ctx) error {
	id := c.Params("id")

	major := new(models.Majors)

	majorInput := new(models.MajorInput)

	if err := c.BodyParser(majorInput); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	connection.DB.Model(&major).Where("id = ?", id).Updates(majorInput)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Major updated",
		"data":    majorInput,
	})
}

func DeleteMajor(c *fiber.Ctx) error {
	id := c.Params("id")
	var major models.Majors

	connection.DB.Find(&major, id)

	if major.Name == "" {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "No major found with ID",
			"data":    nil,
		})
	}

	connection.DB.Delete(&major)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Major deleted",
	})
}
