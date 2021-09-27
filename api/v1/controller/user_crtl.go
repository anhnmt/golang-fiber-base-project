package controller

import (
	"context"
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/xdorro/golang-fiber-base-project/api/v1/repository"
	"github.com/xdorro/golang-fiber-base-project/internal/payload/request"
	"github.com/xdorro/golang-fiber-base-project/pkg/ent"
	"github.com/xdorro/golang-fiber-base-project/util"
)

type UserControllerImpl struct {
	UserRepository repository.UserRepository
}

func NewUserController(ctx context.Context, client *ent.Client) UserController {
	if userController == nil {
		once = &sync.Once{}

		once.Do(func() {
			userController = &UserControllerImpl{
				UserRepository: repository.NewUserRepository(ctx, client),
			}

			log.Println("Create new UserController")
		})
	}

	return userController
}

// FindAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Router /v1/users [get]
func (ctrl *UserControllerImpl) FindAllUsers(*fiber.Ctx) error {
	status := []int{
		util.StatusActive,
		util.StatusDraft,
	}

	users, err := ctrl.UserRepository.FindAllUsersByStatusIn(status)

	if err != nil {
		return util.ResponseBadRequest(err.Error())
	}

	return util.ResponseSuccess("Thành công", users)
}

func (ctrl *UserControllerImpl) CreateNewUsers(c *fiber.Ctx) error {
	userRequest := new(request.UserRequest)

	if err := c.BodyParser(userRequest); err != nil {
		return util.ResponseError(err.Error(), nil)
	}

	hash, _ := util.HashPassword(userRequest.Password)

	user := new(ent.User)
	{
		user.Name = userRequest.Name
		user.Email = userRequest.Email
		user.Password = hash
		user.Status = util.StatusActive
	}

	_, err := ctrl.UserRepository.CreateUser(user)

	if err != nil {
		return util.ResponseBadRequest(err.Error())
	}

	return util.ResponseSuccess("Thành công")
}

func (ctrl *UserControllerImpl) FindUserByID(c *fiber.Ctx) error {
	userID, _ := c.ParamsInt("userId")
	status := []int{
		util.StatusDelete,
	}

	user, err := ctrl.UserRepository.FindUserByIDAndStatusNotInAndIgnorePassword(int64(userID), status)

	if err != nil || user.ID == 0 {
		return util.ResponseBadRequest("User không tồn tại")
	}

	return util.ResponseSuccess("Thành công", user)
}

func (ctrl *UserControllerImpl) UpdateUserByID(c *fiber.Ctx) error {
	userID, _ := c.ParamsInt("userId")
	status := []int{
		util.StatusDelete,
	}

	user, err := ctrl.UserRepository.FindUserByIDAndStatusNotInAndIgnorePassword(int64(userID), status)

	if err != nil || user.ID == 0 {
		return util.ResponseBadRequest("User không tồn tại")
	}

	userRequest := new(request.UserRequest)

	if err = c.BodyParser(userRequest); err != nil {
		return util.ResponseBadRequest(err.Error())
	}

	user.Name = userRequest.Name
	user.Email = userRequest.Email

	_, err = ctrl.UserRepository.UpdateUser(user)

	if err != nil {
		return util.ResponseBadRequest(err.Error())
	}

	return util.ResponseSuccess("Thành công")
}

func (ctrl *UserControllerImpl) DeleteUserByID(c *fiber.Ctx) error {
	userID, _ := c.ParamsInt("userId")
	status := []int{
		util.StatusDelete,
	}

	user, err := ctrl.UserRepository.FindUserByIDAndStatusNotInAndIgnorePassword(int64(userID), status)

	if err != nil || user.ID == 0 {
		return util.ResponseBadRequest("User không tồn tại")
	}

	user.Status = util.StatusDelete

	_, err = ctrl.UserRepository.UpdateStatus(user)

	if err != nil {
		return util.ResponseBadRequest(err.Error())
	}

	return util.ResponseSuccess("Thành công")
}
