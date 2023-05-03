package main

import (
	Db "github.com/ashokdivi/task-2/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	Db.ConnectDb()
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":3030")
}
