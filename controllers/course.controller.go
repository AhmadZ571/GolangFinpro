package controllers

import (
	"github.com/AhmadZ571/GOLANG.git/connection"
	"github.com/AhmadZ571/GOLANG.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllCourse(c *fiber.Ctx) error {
	var courses []models.Courses

	connection.DB.Preload("Majors").Find(&courses)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All courses",
		"data":    courses,
	})
}

func GetCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var courses models.Courses

	connection.DB.Preload("Majors").Find(&courses, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Course found",
		"data":    courses,
	})
}

func CreateCourse(c *fiber.Ctx) error {
	course := new(models.Courses)

	courseInput := new(models.CourseInput)

	if err := c.BodyParser(course); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	connection.DB.Create(&course)

	courseInput.Name = course.Name
	courseInput.Major_Id = course.Major_Id

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Course created",
		"data":    courseInput,
	})
}

func UpdateCourse(c *fiber.Ctx) error {
	id := c.Params("id")

	course := new(models.Courses)

	courseInput := new(models.CourseInput)

	if err := c.BodyParser(courseInput); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	connection.DB.Model(&course).Where("id = ?", id).Updates(courseInput)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Course updated",
		"data":    courseInput,
	})

}

func DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var courses models.Courses

	connection.DB.Find(&courses, id)
	if courses.ID == 0 {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "No course found with given ID",
		})
	}

	connection.DB.Delete(&courses)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Course deleted",
	})
}
