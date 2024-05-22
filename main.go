package main

import (
	"app/app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()      // создание нового приложения
	routes.SetupRoutes(app) //
	app.Listen(":3000")     // слушаем порт 3000

}
