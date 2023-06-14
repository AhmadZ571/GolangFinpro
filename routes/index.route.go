package routes

import (
	"github.com/AhmadZ571/GOLANG.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	// Course API Routes
	app.Get("/api/course", controllers.GetAllCourse)
	app.Get("/api/course/:id", controllers.GetCourse)
	app.Post("/api/course", controllers.CreateCourse)
	app.Put("/api/course/:id", controllers.UpdateCourse)
	app.Delete("/api/course/:id", controllers.DeleteCourse)

	// Lecture API Routes
	app.Get("/api/lecture", controllers.GetAllLecture)
	app.Get("/api/lecture/:id", controllers.GetLecture)
	app.Post("/api/lecture", controllers.CreateLecture)
	app.Put("/api/lecture/:id", controllers.UpdateLecture)
	app.Delete("/api/lecture/:id", controllers.DeleteLecture)

	// Major API Routes
	app.Get("/api/major", controllers.GetAllMajor)
	app.Get("/api/major/:id", controllers.GetMajor)
	app.Post("/api/major", controllers.CreateMajor)
	app.Put("/api/major/:id", controllers.UpdateMajor)
	app.Delete("/api/major/:id", controllers.DeleteMajor)

	// Student API Routes
	app.Get("/api/student", controllers.GetAllStudent)
	app.Get("/api/student/:id", controllers.GetStudent)
	app.Post("/api/student", controllers.CreateStudent)
	app.Put("/api/student/:id", controllers.UpdateStudent)
	app.Delete("/api/student/:id", controllers.DeleteStudent)

	// User API Routes
	app.Get("/api/user", controllers.GetAllUser)
	app.Get("/api/user/:id", controllers.GetUser)
	app.Post("/api/user", controllers.CreateUser)
	app.Put("/api/user/:id", controllers.UpdateUser)
	app.Delete("/api/user/:id", controllers.DeleteUser)
}
