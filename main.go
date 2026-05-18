package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello world")
	})
	app.Listen(":6000")
}
