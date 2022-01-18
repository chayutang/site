package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

// PublicViews func for describe group of public routes.
func PublicViews(a *fiber.App) {

	a.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "ChayuTang",
		})
	})

	// Or extend your config for customization
	a.Use(favicon.New(favicon.Config{
		File: "./public/upload/cropped-favicon-circle-32x32.png",
	}))

	a.Static("/js", "./public/js")
	a.Static("/css", "./public/css")
	a.Static("/sucai", "./public/sucai")
	a.Static("/upload", "./public/upload")
	a.Static("/static", "./public/static")
}
