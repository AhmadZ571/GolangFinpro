package controllers

import (
	"github.com/AhmadZ571/GOLANG.git/connection"
	"github.com/AhmadZ571/GOLANG.git/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllStudent(c *fiber.Ctx) error {
	var student []models.Student

	connection.DB.Preload("Majors").Find(&student)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All student",
		"data":    student,
	})
}

func GetStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student models.Student

	connection.DB.Preload("Majors").Find(&student, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Student found",
		"data":    student,
	})
}

func CreateStudent(c *fiber.Ctx) error {
	student := new(models.Student)

	studentInput := new(models.StudentInput)

	if err := c.BodyParser(student); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	connection.DB.Create(&student)

	studentInput.Name = student.Name
	studentInput.Major_Id = student.Major_Id

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Student created",
		"data":    studentInput,
	})
}

func UpdateStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	student := new(models.Student)

	studentInput := new(models.StudentInput)

	if err := c.BodyParser(studentInput); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"data":    err,
		})
	}

	connection.DB.Model(&student).Where("id = ?", id).Updates(studentInput)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Student updated",
		"data":    studentInput,
	})

}

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student models.Student

	connection.DB.Delete(&student, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Student deleted",
	})
}
