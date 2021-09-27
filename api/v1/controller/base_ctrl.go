package controller

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	once           *sync.Once
	authController *AuthControllerImpl
	userController *UserControllerImpl
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	CurrentUser(c *fiber.Ctx) error
	Restricted(c *fiber.Ctx) error
}

type UserController interface {
	FindAllUsers(c *fiber.Ctx) error
	FindUserByID(c *fiber.Ctx) error
	CreateNewUsers(c *fiber.Ctx) error
	UpdateUserByID(c *fiber.Ctx) error
	DeleteUserByID(c *fiber.Ctx) error
}
