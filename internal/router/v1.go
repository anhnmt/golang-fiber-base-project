package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/xdorro/golang-fiber-base-project/api/v1/controller"
	"github.com/xdorro/golang-fiber-base-project/config"
	"github.com/xdorro/golang-fiber-base-project/internal/middleware"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent"
	"github.com/xdorro/golang-fiber-base-project/util"
)

func V1Router(app fiber.Router, conf *config.DefaultConfig, client *ent.Client) {
	v1 := app.Group("/v1")
	{
		v1.Get("/", func(c *fiber.Ctx) error {
			return util.ResponseSuccess("API V1")
		})

		auth := v1.Group("/auth")
		{
			authController := controller.NewAuthController(client, conf.Ctx)
			auth.Post("/login", authController.Login)

			auth.Use(middleware.Protected())
			auth.Get("/current_user", authController.CurrentUser)
			auth.Get("/restricted", authController.Restricted)
		}

		user := v1.Group("/users")
		{
			userController := controller.NewUserController(client, conf.Ctx)
			user.Get("/", userController.FindAllUsers)
			user.Post("/", userController.CreateNewUsers)
			user.Get("/:userId", userController.FindUserById)
			user.Put("/:userId", userController.UpdateUserById)
			user.Delete("/:userId", userController.DeleteUserById)
		}
	}
}
