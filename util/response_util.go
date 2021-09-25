package util

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/internal/payload/response"
)

// ResponseSuccess : returning json structure for success request
func ResponseSuccess(message string, data ...interface{}) error {
	return JsonResponse(fiber.StatusOK, message, data)
}

// ResponseNotFound : returning json structure for notfound request
func ResponseNotFound(message string) error {
	return JsonResponse(fiber.StatusNotFound, message)
}

// ResponseError : returning json structure for error request
func ResponseError(message string, data ...interface{}) error {
	return JsonResponse(fiber.StatusInternalServerError, message, data)
}

// ResponseUnauthorized : returning json structure for validator error request
func ResponseUnauthorized(message string, data ...interface{}) error {
	return JsonResponse(fiber.StatusUnauthorized, message, data)
}

// ResponseBadRequest : returning json structure for validator error request
func ResponseBadRequest(message string, data ...interface{}) error {
	return JsonResponse(fiber.StatusBadRequest, message, data)
}

func JsonResponse(status int, message string, data ...interface{}) error {
	msg, _ := json.Marshal(&response.Response{
		Status:  status,
		Message: message,
		Data:    GetFirst(data),
	})

	return fiber.NewError(status, string(msg))
}
