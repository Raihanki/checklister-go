package handlers

import (
	"errors"
	"strconv"

	"github.com/Raihanki/checklisters/internal/dto"
	"github.com/Raihanki/checklisters/internal/errpkg"
	"github.com/Raihanki/checklisters/internal/helpers"
	"github.com/Raihanki/checklisters/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ChecklistItemHandler struct {
	ChecklistItemService services.ChecklistItemService
}

func (h ChecklistItemHandler) Store(ctx *fiber.Ctx) error {
	checklistIdStr := ctx.Params("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Id", nil)
	}

	request := dto.ChecklistItemRequest{}
	err = ctx.BodyParser(&request)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Request Body", nil)
	}

	errValidation := helpers.ValidateStruct(request)
	if errValidation != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Request Body", errValidation)
	}

	request.ChecklistId = checklistId
	err = h.ChecklistItemService.Create(ctx, request)
	if err != nil {
		if errors.Is(err, errpkg.ErrChecklistItemNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Not Found", nil)
		}
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 201, "Checklist Item Created", nil)
}

func (h ChecklistItemHandler) Index(ctx *fiber.Ctx) error {
	checklistIdStr := ctx.Params("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Id", nil)
	}

	response, err := h.ChecklistItemService.GetAll(ctx, checklistId)
	if err != nil {
		if errors.Is(err, errpkg.ErrChecklistItemNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Not Found", nil)
		}
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 200, "Checklist Items", response)
}

func (h ChecklistItemHandler) Destroy(ctx *fiber.Ctx) error {
	checklistIdStr := ctx.Params("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Id", nil)
	}

	checklistItemIdStr := ctx.Params("checklistItemId")
	checklistItemId, err := strconv.Atoi(checklistItemIdStr)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Item Id", nil)
	}

	err = h.ChecklistItemService.Delete(ctx, checklistId, checklistItemId)
	if err != nil {
		if errors.Is(err, errpkg.ErrChecklistItemNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Item Not Found", nil)
		}
		if errors.Is(err, errpkg.ErrChecklistNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Not Found", nil)
		}
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 204, "Checklist Item Deleted", nil)
}

func (h ChecklistItemHandler) Complete(ctx *fiber.Ctx) error {
	checklistIdStr := ctx.Params("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Id", nil)
	}

	checklistItemIdStr := ctx.Params("checklistItemId")
	checklistItemId, err := strconv.Atoi(checklistItemIdStr)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Item Id", nil)
	}

	response, err := h.ChecklistItemService.Complete(ctx, checklistItemId, checklistId)
	if err != nil {
		if errors.Is(err, errpkg.ErrChecklistItemNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Item Not Found", nil)
		}
		if errors.Is(err, errpkg.ErrChecklistNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Not Found", nil)
		}
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 200, "Checklist Item Completed", response)
}

func (h ChecklistItemHandler) Rename(ctx *fiber.Ctx) error {
	checklistIdStr := ctx.Params("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Id", nil)
	}

	checklistItemIdStr := ctx.Params("checklistItemId")
	checklistItemId, err := strconv.Atoi(checklistItemIdStr)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Item Id", nil)
	}

	request := dto.UpdateChecklistItemRequest{}
	err = ctx.BodyParser(&request)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Request Body", nil)
	}

	errValidation := helpers.ValidateStruct(request)
	if errValidation != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Request Body", errValidation)
	}

	request.ChecklistId = checklistId
	response, err := h.ChecklistItemService.Rename(ctx, request, checklistItemId)
	if err != nil {
		if errors.Is(err, errpkg.ErrChecklistItemNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Item Not Found", nil)
		}
		if errors.Is(err, errpkg.ErrChecklistNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Not Found", nil)
		}
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 200, "Checklist Item Updated", response)
}

func (h ChecklistItemHandler) Show(ctx *fiber.Ctx) error {
	checklistIdStr := ctx.Params("checklistId")
	checklistId, err := strconv.Atoi(checklistIdStr)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Id", nil)
	}

	checklistItemIdStr := ctx.Params("checklistItemId")
	checklistItemId, err := strconv.Atoi(checklistItemIdStr)
	if err != nil {
		return helpers.JsonResponse(ctx, 400, "Invalid Checklist Item Id", nil)
	}

	response, err := h.ChecklistItemService.GetById(ctx, checklistId, checklistItemId)
	if err != nil {
		if errors.Is(err, errpkg.ErrChecklistItemNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Item Not Found", nil)
		}
		if errors.Is(err, errpkg.ErrChecklistNotFound) {
			return helpers.JsonResponse(ctx, 404, "Checklist Not Found", nil)
		}
		return helpers.JsonResponse(ctx, 500, "", nil)
	}

	return helpers.JsonResponse(ctx, 200, "Checklist Item", response)
}
