package main

import (
	"github.com/AhmadZ571/GOLANG.git/connection"
	"github.com/AhmadZ571/GOLANG.git/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	connection.Connect()

	app := fiber.New()

	routes.Init(app)

	app.Listen(":8080")
}
