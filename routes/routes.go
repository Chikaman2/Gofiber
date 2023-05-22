package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/chikaman2/Gofiber/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)

}
