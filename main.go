package main

import (
	"github.com/chikaman2/Gofiber/database"
	"github.com/chikaman2/Gofiber/routes"
	"github.com/gofiber/fiber/v2"
)

const portNumber = ":3000"

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(portNumber)

}
