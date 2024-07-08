package routers

import "github.com/gofiber/fiber/v2"

func ExampleRoute(app *fiber.App) {
	route := app.Group("/")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("OK")
	})
}
