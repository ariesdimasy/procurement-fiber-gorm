package routers

import (
	"procurement-fiber-gorm/controllers"

	"github.com/gofiber/fiber/v2"
)

func ItemRouter(app *fiber.App) {
	router := app.Group("/items")
	router.Get("/", controllers.GetItems)
	router.Post("/", controllers.CreateItem)
}
