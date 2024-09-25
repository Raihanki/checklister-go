package services

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/Raihanki/checklisters/internal/domain"
	"github.com/Raihanki/checklisters/internal/dto"
	"github.com/Raihanki/checklisters/internal/errpkg"
	"github.com/Raihanki/checklisters/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

type ChecklistItemService interface {
	GetAll(c *fiber.Ctx, checklistId int) ([]dto.ChecklistItemResponse, error)
	GetById(c *fiber.Ctx, checklistId int, checklistItemId int) (dto.ChecklistItemResponse, error)
	Create(c *fiber.Ctx, request dto.ChecklistItemRequest) error
	Complete(c *fiber.Ctx, checklistItemId int, checklistId int) (dto.ChecklistItemResponse, error)
	Rename(c *fiber.Ctx, request dto.UpdateChecklistItemRequest, checklistItemId int) (dto.ChecklistItemResponse, error)
	Delete(c *fiber.Ctx, checklistId int, checklistItemId int) error
}

type ChecklistItemServiceImpl struct {
	DB                      *sql.DB
	ChecklistItemRepository repositories.ChecklistItemRepository
	ChecklistRepository     repositories.ChecklistRepository
}

func NewChecklistItemService(db *sql.DB, checklistItemRepo repositories.ChecklistItemRepository, checklistRepo repositories.ChecklistRepository) ChecklistItemService {
	return &ChecklistItemServiceImpl{
		DB:                      db,
		ChecklistItemRepository: checklistItemRepo,
		ChecklistRepository:     checklistRepo,
	}
}

func (s *ChecklistItemServiceImpl) GetAll(c *fiber.Ctx, checklistId int) ([]dto.ChecklistItemResponse, error) {
	userIdStr := c.Locals("user_id").(string)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("error convert user_id to int :: %v", err)
		return nil, err
	}

	_, err = s.ChecklistRepository.GetById(c, s.DB, checklistId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errpkg.ErrChecklistNotFound
		}
		log.Printf("error when get checklist by id - %v", err)
		return nil, err
	}

	checklistItems, err := s.ChecklistItemRepository.GetAll(c, s.DB, checklistId)
	if err != nil {
		log.Printf("error when get all checklist items - %v", err)
		return nil, err
	}

	var resChecklistItems []dto.ChecklistItemResponse
	for _, item := range checklistItems {
		resChecklistItems = append(resChecklistItems, dto.ChecklistItemResponse{
			Id:          item.Id,
			ChecklistId: item.ChecklistId,
			ItemName:    item.ItemName,
			CompletedAt: item.CompletedAt,
			CreatedAt:   *item.CretedAt,
		})
	}

	return resChecklistItems, nil
}

func (s *ChecklistItemServiceImpl) GetById(c *fiber.Ctx, checklistId int, checklistItemId int) (dto.ChecklistItemResponse, error) {
	userIdStr := c.Locals("user_id").(string)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("error convert user_id to int :: %v", err)
		return dto.ChecklistItemResponse{}, err
	}

	_, err = s.ChecklistRepository.GetById(c, s.DB, checklistId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.ChecklistItemResponse{}, errpkg.ErrChecklistNotFound
		}
		log.Printf("error when get checklist by id - %v", err)
		return dto.ChecklistItemResponse{}, err
	}

	checklistItem, err := s.ChecklistItemRepository.GetById(c, s.DB, checklistId, checklistItemId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.ChecklistItemResponse{}, errpkg.ErrChecklistItemNotFound
		}
		log.Printf("error when get checklist item by id - %v", err)
		return dto.ChecklistItemResponse{}, err
	}

	return dto.ChecklistItemResponse{
		Id:          checklistItem.Id,
		ItemName:    checklistItem.ItemName,
		ChecklistId: checklistItem.ChecklistId,
		CompletedAt: checklistItem.CompletedAt,
		CreatedAt:   *checklistItem.CretedAt,
	}, nil
}

func (s *ChecklistItemServiceImpl) Create(c *fiber.Ctx, request dto.ChecklistItemRequest) error {
	userIdStr := c.Locals("user_id").(string)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("error convert user_id to int :: %v", err)
		return err
	}

	_, err = s.ChecklistRepository.GetById(c, s.DB, request.ChecklistId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errpkg.ErrChecklistNotFound
		}
		log.Printf("error when get checklist by id - %v", err)
		return err
	}

	checklistItem := domain.ChecklistItem{
		ItemName:    request.ItemName,
		ChecklistId: request.ChecklistId,
	}

	err = s.ChecklistItemRepository.Create(c, s.DB, checklistItem)
	if err != nil {
		return err
	}

	return nil
}

func (s *ChecklistItemServiceImpl) Complete(c *fiber.Ctx, checklistItemId int, checklistId int) (dto.ChecklistItemResponse, error) {
	userIdStr := c.Locals("user_id").(string)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("error convert user_id to int :: %v", err)
		return dto.ChecklistItemResponse{}, err
	}

	_, err = s.ChecklistRepository.GetById(c, s.DB, checklistId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.ChecklistItemResponse{}, errpkg.ErrChecklistNotFound
		}
		log.Printf("error when get checklist by id - %v", err)
		return dto.ChecklistItemResponse{}, err
	}

	checklistItem, err := s.ChecklistItemRepository.Complete(c, s.DB, checklistItemId, checklistId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.ChecklistItemResponse{}, errpkg.ErrChecklistItemNotFound
		}
		log.Printf("error when complete checklist item - %v", err)
		return dto.ChecklistItemResponse{}, err
	}

	return dto.ChecklistItemResponse{
		Id:          checklistItem.Id,
		ItemName:    checklistItem.ItemName,
		ChecklistId: checklistItem.ChecklistId,
		CompletedAt: checklistItem.CompletedAt,
		CreatedAt:   *checklistItem.CretedAt,
	}, nil
}

func (s *ChecklistItemServiceImpl) Rename(c *fiber.Ctx, request dto.UpdateChecklistItemRequest, checklistItemId int) (dto.ChecklistItemResponse, error) {
	userIdStr := c.Locals("user_id").(string)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("error convert user_id to int :: %v", err)
		return dto.ChecklistItemResponse{}, err
	}

	_, err = s.ChecklistRepository.GetById(c, s.DB, request.ChecklistId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.ChecklistItemResponse{}, errpkg.ErrChecklistNotFound
		}
		log.Printf("error when get checklist by id - %v", err)
		return dto.ChecklistItemResponse{}, err
	}

	checklistItem := domain.ChecklistItem{
		ItemName:    request.ItemName,
		ChecklistId: request.ChecklistId,
	}

	checklistItem, err = s.ChecklistItemRepository.Rename(c, s.DB, checklistItem, checklistItemId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.ChecklistItemResponse{}, errpkg.ErrChecklistItemNotFound
		}
		log.Printf("error when rename checklist item - %v", err)
		return dto.ChecklistItemResponse{}, err
	}

	return dto.ChecklistItemResponse{
		Id:          checklistItem.Id,
		ItemName:    checklistItem.ItemName,
		CompletedAt: checklistItem.CompletedAt,
		CreatedAt:   *checklistItem.CretedAt,
	}, nil
}

func (s *ChecklistItemServiceImpl) Delete(c *fiber.Ctx, checklistId int, checklistItemId int) error {
	userIdStr := c.Locals("user_id").(string)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("error convert user_id to int :: %v", err)
		return err
	}

	_, err = s.ChecklistRepository.GetById(c, s.DB, checklistId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errpkg.ErrChecklistNotFound
		}
		log.Printf("error when get checklist by id - %v", err)
		return err
	}

	err = s.ChecklistItemRepository.Delete(c, s.DB, checklistId, checklistItemId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errpkg.ErrChecklistItemNotFound
		}
		log.Printf("error when delete checklist item - %v", err)
		return err
	}

	return nil
}
