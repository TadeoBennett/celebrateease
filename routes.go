package main

import (
	"tadeobennett/celebrateease/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.PostsIndex)

    // Custom 404 handler
    app.Use(func(c *fiber.Ctx) error {
        // Respond with a 404 status code
        return c.Status(fiber.StatusNotFound).SendString("404 Not Found - Redirecting to home page")
    })
}
