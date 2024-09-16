package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello word")
	})
	app.Listen(":3000")
}
