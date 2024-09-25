package handlers

import (
	"errors"
	"net/http"

	"github.com/Raihanki/checklisters/internal/dto"
	"github.com/Raihanki/checklisters/internal/errpkg"
	"github.com/Raihanki/checklisters/internal/helpers"
	"github.com/Raihanki/checklisters/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService services.UserService
}

func (h UserHandler) Register(ctx *fiber.Ctx) error {
	request := dto.RegisterRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Request Body", nil)
	}

	errValidation := helpers.ValidateStruct(request)
	if errValidation != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Request Body", errValidation)
	}

	response, err := h.UserService.Register(ctx, request)
	if err != nil {
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 201, http.StatusText(http.StatusCreated), response)
}

func (h UserHandler) Login(ctx *fiber.Ctx) error {
	request := dto.LoginRequest{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Request Body", nil)
	}

	errValidation := helpers.ValidateStruct(request)
	if errValidation != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Request Body", errValidation)
	}

	response, err := h.UserService.Login(ctx, request)
	if err != nil {
		if errors.Is(err, errpkg.ErrInvalidEmailOrPassword) {
			return helpers.JsonResponse(ctx, 400, "Invalid Username or Password", nil)
		}
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 200, http.StatusText(http.StatusOK), response)
}
