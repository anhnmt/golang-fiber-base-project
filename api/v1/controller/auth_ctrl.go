package controller

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/xdorro/golang-fiber-base-project/api/v1/repository"
	"github.com/xdorro/golang-fiber-base-project/config"
	"github.com/xdorro/golang-fiber-base-project/internal/payload/request"
	"github.com/xdorro/golang-fiber-base-project/internal/payload/response"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent"
	"github.com/xdorro/golang-fiber-base-project/util"
)

type AuthControllerImpl struct {
	UserRepository repository.UserRepository
}

func NewAuthController(ctx context.Context, client *ent.Client) AuthController {
	if authController == nil {
		once = &sync.Once{}

		once.Do(func() {
			authController = &AuthControllerImpl{
				UserRepository: repository.NewUserRepository(ctx, client),
			}

			log.Println("Create new AuthController")
		})
	}

	return authController
}

func (ctrl *AuthControllerImpl) Login(c *fiber.Ctx) error {
	var loginRequest request.LoginRequest
	status := []int{
		util.StatusDraft,
		util.StatusDelete,
	}

	if err := c.BodyParser(&loginRequest); err != nil {
		return util.ResponseBadRequest(err.Error())
	}

	user, err := ctrl.UserRepository.FindUserByEmailAndStatusNotIn(loginRequest.Email, status)

	if err != nil || user.ID == 0 {
		return util.ResponseBadRequest("User không tồn tại", err)
	}

	log.Println(user)

	if !util.CheckPasswordHash(loginRequest.Password, user.Password) {
		return util.ResponseBadRequest("Mật khẩu không chính xác")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	accessToken, err := token.SignedString([]byte(config.Config().JWTSecret))
	if err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	result := &response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: accessToken,
	}

	return util.ResponseSuccess("Thành công", result)
}

func (ctrl *AuthControllerImpl) CurrentUser(c *fiber.Ctx) error {
	claims, err := util.GetTokenClaims(c)
	if err != nil {
		return err
	}

	ID := int64(claims["id"].(float64))
	status := []int{
		util.StatusDraft,
		util.StatusDelete,
	}

	user, err := ctrl.UserRepository.FindUserByIDAndStatusNotInAndIgnorePassword(ID, status)

	if err != nil || user.ID == 0 {
		return util.ResponseBadRequest("User không tồn tại", err)
	}

	return util.ResponseSuccess("Thành công", user)
}

func (ctrl *AuthControllerImpl) Restricted(c *fiber.Ctx) error {
	claims, err := util.GetTokenClaims(c)
	if err != nil {
		return err
	}

	email := claims["email"].(string)

	return util.ResponseSuccess("Welcome " + email)
}
