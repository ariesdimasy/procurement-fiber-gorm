package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"procurement-fiber-gorm/config"
	"procurement-fiber-gorm/routers"
)

func main() {

	config.DatabaseConnection()

	app := fiber.New()

	routers.ItemRouter(app)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"foo": "bar",
		})
	})

	log.Fatal(app.Listen(":5760"))
}
