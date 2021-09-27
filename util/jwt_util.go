package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// GetTokenClaims get token claims
func GetTokenClaims(c *fiber.Ctx) (jwt.MapClaims, error) {
	local := c.Locals("user")

	if local == nil {
		return nil, ResponseBadRequest("User chưa đăng nhập")
	}

	token := local.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}

// ValidToken check valid token
func ValidToken(token *jwt.Token, id int64) bool {
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["id"].(int64)

	return id != userID
}

// HashPassword hash password
func HashPassword(password string) (string, error) {
	const COST = 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), COST)
	return string(bytes), err
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
