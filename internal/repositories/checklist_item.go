package repositories

import (
	"database/sql"

	"github.com/Raihanki/checklisters/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type ChecklistItemRepository interface {
	GetAll(c *fiber.Ctx, db *sql.DB, checklistId int) ([]domain.ChecklistItem, error)
	GetById(c *fiber.Ctx, db *sql.DB, checklistId int, checklistItemId int) (domain.ChecklistItem, error)
	Create(c *fiber.Ctx, db *sql.DB, checklistItem domain.ChecklistItem) error
	Complete(c *fiber.Ctx, db *sql.DB, checklistItemId int, checklistId int) (domain.ChecklistItem, error)
	Rename(c *fiber.Ctx, db *sql.DB, checklistItem domain.ChecklistItem, checklistItemId int) (domain.ChecklistItem, error)
	Delete(c *fiber.Ctx, db *sql.DB, checklistId int, checklistItemId int) error
}

type ChecklistItemRepositoryImpl struct {
}

func NewChecklistItemRepository() ChecklistItemRepository {
	return &ChecklistItemRepositoryImpl{}
}

func (repo *ChecklistItemRepositoryImpl) GetAll(c *fiber.Ctx, db *sql.DB, checklistId int) ([]domain.ChecklistItem, error) {
	query := "SELECT id, item_name, checklist_id, completed_at, created_at FROM checklist_items WHERE checklist_id = $1"
	rows, err := db.QueryContext(c.Context(), query, checklistId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var checklistItems []domain.ChecklistItem
	for rows.Next() {
		checklistItem := domain.ChecklistItem{}
		err := rows.Scan(&checklistItem.Id, &checklistItem.ItemName, &checklistItem.ChecklistId, &checklistItem.CompletedAt, &checklistItem.CretedAt)
		if err != nil {
			return nil, err
		}
		checklistItems = append(checklistItems, checklistItem)
	}

	return checklistItems, nil
}

func (repo *ChecklistItemRepositoryImpl) GetById(c *fiber.Ctx, db *sql.DB, checklistId int, checklistItemId int) (domain.ChecklistItem, error) {
	query := "SELECT id, item_name, checklist_id, completed_at, created_at FROM checklist_items WHERE id = $1 AND checklist_id = $2"
	row := db.QueryRowContext(c.Context(), query, checklistItemId, checklistId)
	checklist := domain.ChecklistItem{}
	err := row.Scan(&checklist.Id, &checklist.ItemName, &checklist.ChecklistId, &checklist.CompletedAt, &checklist.CretedAt)
	if err != nil {
		return domain.ChecklistItem{}, err
	}

	return checklist, nil
}

func (repo *ChecklistItemRepositoryImpl) Create(c *fiber.Ctx, db *sql.DB, checklistItem domain.ChecklistItem) error {
	query := "INSERT INTO checklist_items (item_name, checklist_id) VALUES ($1, $2)"
	_, err := db.ExecContext(c.Context(), query, checklistItem.ItemName, checklistItem.ChecklistId)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ChecklistItemRepositoryImpl) Complete(c *fiber.Ctx, db *sql.DB, checklistItemId int, checklistId int) (domain.ChecklistItem, error) {
	query := "UPDATE checklist_items SET completed_at = NOW() WHERE id = $1 AND checklist_id = $2 RETURNING id, item_name, checklist_id, completed_at, created_at"
	row := db.QueryRowContext(c.Context(), query, checklistItemId, checklistId)
	checklist := domain.ChecklistItem{}
	err := row.Scan(&checklist.Id, &checklist.ItemName, &checklist.ChecklistId, &checklist.CompletedAt, &checklist.CretedAt)
	if err != nil {
		return domain.ChecklistItem{}, err
	}

	return checklist, nil
}

func (repo *ChecklistItemRepositoryImpl) Rename(c *fiber.Ctx, db *sql.DB, checklistItem domain.ChecklistItem, checklistItemId int) (domain.ChecklistItem, error) {
	query := "UPDATE checklist_items SET item_name = $1 WHERE id = $2 AND checklist_id = $3 RETURNING id, item_name, checklist_id, completed_at, created_at"
	row := db.QueryRowContext(c.Context(), query, checklistItem.ItemName, checklistItemId, checklistItem.ChecklistId)
	checklist := domain.ChecklistItem{}
	err := row.Scan(&checklist.Id, &checklist.ItemName, &checklist.ChecklistId, &checklist.CompletedAt, &checklist.CretedAt)
	if err != nil {
		return domain.ChecklistItem{}, err
	}

	return checklist, nil
}

func (repo *ChecklistItemRepositoryImpl) Delete(c *fiber.Ctx, db *sql.DB, checklistId int, checklistItemId int) error {
	query := "DELETE FROM checklist_items WHERE id = $1 AND checklist_id = $2"
	row, err := db.ExecContext(c.Context(), query, checklistItemId, checklistId)
	if r, _ := row.RowsAffected(); r == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}

	return nil
}
