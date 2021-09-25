package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"github.com/xdorro/golang-fiber-base-project/config"
	"github.com/xdorro/golang-fiber-base-project/internal/database"
	"github.com/xdorro/golang-fiber-base-project/internal/middleware"
	"github.com/xdorro/golang-fiber-base-project/internal/router"
)

// @title Golang Fiber Simple Project
// @version 1.0.0
// @description This is a base project for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name Nguyen Manh Tuan Anh
// @contact.email xdorro@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	conf := config.Config()

	app := fiber.New(fiber.Config{
		AppName: conf.AppName,

		Prefork: conf.AppPrefork,

		// Override default error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Default 500 statusCode
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				// Override status code if fiber.Error type
				code = e.Code
			}
			// Set Content-Type: application/json; charset=utf-8
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

			// Return statusCode with error message
			return c.Status(code).Send([]byte(err.Error()))
		},
	})

	// Connect DB
	client := database.Connection(conf)

	// Config Middleware
	middleware.Middleware(app, conf)

	// Config Router
	router.Router(app, conf, client)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		log.Printf("Close database connection")
		_ = database.Close()

		log.Printf("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	if err := app.Listen(fmt.Sprintf(":%d", conf.AppPort)); err != nil {
		log.Printf("Oops... server is not running! error: %v", err)
	}
}
