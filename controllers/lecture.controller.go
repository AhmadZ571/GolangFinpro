package controllers

import (
	"github.com/AhmadZ571/GOLANG.git/connection"
	"github.com/AhmadZ571/GOLANG.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllLecture(c *fiber.Ctx) error {
	var lecture []models.Lecture

	connection.DB.Preload("Courses").Preload("Courses.Majors").Find(&lecture)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All lecture",
		"data":    lecture,
	})

}

func GetLecture(c *fiber.Ctx) error {
	id := c.Params("id")
	var lecture models.Lecture

	connection.DB.Preload("Courses").Preload("Courses.Majors").Find(&lecture, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Lecture found",
		"data":    lecture,
	})
}

func CreateLecture(c *fiber.Ctx) error {
	lecture := new(models.Lecture)

	lectureInput := new(models.LectureInput)

	if err := c.BodyParser(lecture); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	connection.DB.Create(&lecture)

	lectureInput.Name = lecture.Name
	lectureInput.Course_Id = lecture.Course_Id

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Lecture created",
		"data":    lectureInput,
	})
}

func UpdateLecture(c *fiber.Ctx) error {
	id := c.Params("id")

	lecture := new(models.Lecture)

	lectureInput := new(models.LectureInput)

	if err := c.BodyParser(lectureInput); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	connection.DB.Model(&lecture).Where("id = ?", id).Updates(lectureInput)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Lecture updated",
		"data":    lectureInput,
	})
}

func DeleteLecture(c *fiber.Ctx) error {
	id := c.Params("id")
	var lecture models.Lecture

	connection.DB.First(&lecture, id)

	if lecture.Name == "" {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "No lecture found with ID",
			"data":    nil,
		})
	}

	connection.DB.Delete(&lecture)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Lecture deleted",
	})
}
