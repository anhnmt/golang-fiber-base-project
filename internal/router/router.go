package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/config"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent"
	"github.com/xdorro/golang-fiber-base-project/util"
)

func Router(app *fiber.App, conf *config.DefaultConfig, client *ent.Client) {
	app.Get("/", func(c *fiber.Ctx) error {
		return util.ResponseSuccess("Welcome to Fiber Go API!")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return util.ResponseSuccess("pong")
	})

	// API Group
	//api := app.Group("/api")
	//{
	//}

	// 404 error
	NotFoundRouter(app)
}

// NotFoundRouter Handler page not found 404
func NotFoundRouter(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return util.ResponseNotFound("Not Found")
		},
	)
}
