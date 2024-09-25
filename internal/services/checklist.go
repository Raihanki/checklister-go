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

type ChecklistService interface {
	GetAll(ctx *fiber.Ctx) ([]dto.ChecklistResponse, error)
	Create(ctx *fiber.Ctx, request dto.CreateChecklist) error
	Delete(ctx *fiber.Ctx, checklistId int) error
}

type ChecklistServiceImpl struct {
	DB                  *sql.DB
	ChecklistRepository repositories.ChecklistRepository
}

func NewChecklistService(repo repositories.ChecklistRepository, db *sql.DB) ChecklistService {
	return &ChecklistServiceImpl{ChecklistRepository: repo, DB: db}
}

func (s *ChecklistServiceImpl) GetAll(ctx *fiber.Ctx) ([]dto.ChecklistResponse, error) {
	userId, err := strconv.Atoi(ctx.Locals("user_id").(string))
	if err != nil {
		log.Printf("error convert user_id to int :: %v", err)
		return nil, err
	}
	checklists, err := s.ChecklistRepository.GetAll(ctx, s.DB, userId)
	if err != nil {
		log.Printf("error when get all checklist - %v", err)
		return nil, err
	}

	var resChecklist []dto.ChecklistResponse
	for _, checklist := range checklists {
		resChecklist = append(resChecklist, dto.ChecklistResponse{
			Id:        checklist.Id,
			Name:      checklist.Name,
			CreatedAt: checklist.CreatedAt,
		})
	}
	return resChecklist, nil
}

func (s *ChecklistServiceImpl) Create(ctx *fiber.Ctx, request dto.CreateChecklist) error {
	tx, err := s.DB.Begin()
	if err != nil {
		log.Printf("error whlie begin transaction in create checklist - %v", err)
		return err
	}
	defer tx.Rollback()

	checklist := domain.Checklist{
		Name: request.Name,
	}
	userId, err := strconv.Atoi(ctx.Locals("user_id").(string))
	if err != nil {
		log.Printf("error convert user_id to int :: %v", err)
		return err
	}
	err = s.ChecklistRepository.Create(ctx, tx, checklist, userId)
	if err != nil {
		log.Printf("error whlie create checklist - %v", err)
		return err
	}

	return nil
}

func (s *ChecklistServiceImpl) Delete(ctx *fiber.Ctx, checklistId int) error {
	tx, err := s.DB.Begin()
	if err != nil {
		log.Printf("error whlie begin transaction in create checklist - %v", err)
		return err
	}
	defer tx.Rollback()

	userId, err := strconv.Atoi(ctx.Locals("user_id").(string))
	if err != nil {
		log.Printf("error convert user_id to int :: %v", err)
		return err
	}
	err = s.ChecklistRepository.Delete(ctx, tx, checklistId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errpkg.ErrChecklistNotFound
		}
		log.Printf("error whlie create checklist - %v", err)
		return err
	}

	return nil
}
