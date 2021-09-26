package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"

	"github.com/xdorro/golang-fiber-base-project/config"
	"github.com/xdorro/golang-fiber-base-project/util"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Config().JWTSecret),

		ErrorHandler: func(_ *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return util.ResponseBadRequest("Missing or malformed JWT")
			}

			return util.ResponseUnauthorized("Invalid or expired JWT")
		},
	})
}
