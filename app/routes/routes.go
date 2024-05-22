// В этой директории находятся подкаталоги для маршрутов

package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error { // определяем http маршрут
		return c.SendString("Hello, Fiber!")
	})
}
