package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/xdorro/golang-fiber-base-project/config"
)

func Middleware(app *fiber.App, conf *config.DefaultConfig) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// Add Recover
	app.Use(recover.New())

	// Add Icon
	app.Use(favicon.New())

	// Add CORS to each route.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowCredentials: true,
	}))

	// Add simple logger.
	if conf.LogEnabled {
		app.Use(logger.New())
	}

	// Add caching.
	// cache.New(),
}
