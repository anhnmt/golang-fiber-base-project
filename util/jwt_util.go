package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetTokenClaims(c *fiber.Ctx) (jwt.MapClaims, error) {
	local := c.Locals("user")

	if local == nil {
		return nil, ResponseBadRequest("User chưa đăng nhập")
	}

	token := local.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}

func ValidToken(token *jwt.Token, ID int64) bool {
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(int64)

	if ID != userId {
		return false
	}

	return true
}

// HashPassword hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
