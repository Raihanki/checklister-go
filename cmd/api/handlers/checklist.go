package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Raihanki/checklisters/internal/dto"
	"github.com/Raihanki/checklisters/internal/errpkg"
	"github.com/Raihanki/checklisters/internal/helpers"
	"github.com/Raihanki/checklisters/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ChecklistHandler struct {
	ChecklistService services.ChecklistService
}

func (h *ChecklistHandler) Index(ctx *fiber.Ctx) error {
	response, err := h.ChecklistService.GetAll(ctx)
	if err != nil {
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 200, http.StatusText(http.StatusOK), response)
}

func (h *ChecklistHandler) Store(ctx *fiber.Ctx) error {
	request := dto.CreateChecklist{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Request Body", nil)
	}

	errValidation := helpers.ValidateStruct(request)
	if errValidation != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Request Body", errValidation)
	}

	err = h.ChecklistService.Create(ctx, request)
	if err != nil {
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 201, http.StatusText(http.StatusCreated), nil)
}

func (h *ChecklistHandler) Destroy(ctx *fiber.Ctx) error {
	strChecklistId := ctx.Params("checklistId")
	checklistId, err := strconv.Atoi(strChecklistId)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Id", nil)
	}

	err = h.ChecklistService.Delete(ctx, checklistId)
	if err != nil {
		if errors.Is(err, errpkg.ErrChecklistNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Not Found", nil)
		}
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 204, http.StatusText(http.StatusNoContent), nil)
}
