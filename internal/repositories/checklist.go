package repositories

import (
	"database/sql"

	"github.com/Raihanki/checklisters/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type ChecklistRepository interface {
	GetAll(c *fiber.Ctx, db *sql.DB, userId int) ([]domain.Checklist, error)
	Create(c *fiber.Ctx, tx *sql.Tx, checklist domain.Checklist, userId int) error
	Delete(c *fiber.Ctx, tx *sql.Tx, checklistId int, userId int) error
	GetById(c *fiber.Ctx, db *sql.DB, checklistId int, userId int) (domain.Checklist, error)
}

type ChecklistRepositoryImpl struct {
}

func NewChecklistRepository() ChecklistRepository {
	return &ChecklistRepositoryImpl{}
}

func (repo ChecklistRepositoryImpl) GetAll(c *fiber.Ctx, db *sql.DB, userId int) ([]domain.Checklist, error) {
	query := "SELECT id, name, created_at FROM checklists WHERE user_id = $1"
	rows, err := db.QueryContext(c.Context(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var checklists []domain.Checklist
	for rows.Next() {
		checklist := domain.Checklist{}
		err := rows.Scan(&checklist.Id, &checklist.Name, &checklist.CreatedAt)
		if err != nil {
			return nil, err
		}
		checklists = append(checklists, checklist)
	}

	return checklists, nil
}

func (repo ChecklistRepositoryImpl) Create(c *fiber.Ctx, tx *sql.Tx, checklist domain.Checklist, userId int) error {
	query := "INSERT INTO checklists (name, user_id) VALUES ($1, $2)"
	_, err := tx.ExecContext(c.Context(), query, checklist.Name, userId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo ChecklistRepositoryImpl) Delete(c *fiber.Ctx, tx *sql.Tx, checklistId int, userId int) error {
	query := "DELETE FROM checklists WHERE id = $1 AND user_id = $2"
	row, err := tx.ExecContext(c.Context(), query, checklistId, userId)
	if r, _ := row.RowsAffected(); r == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo ChecklistRepositoryImpl) GetById(c *fiber.Ctx, db *sql.DB, checklistId int, userId int) (domain.Checklist, error) {
	query := "SELECT id, name, created_at FROM checklists WHERE id = $1 AND user_id = $2"
	row := db.QueryRowContext(c.Context(), query, checklistId, userId)
	checklist := domain.Checklist{}
	err := row.Scan(&checklist.Id, &checklist.Name, &checklist.CreatedAt)
	if err != nil {
		return domain.Checklist{}, err
	}

	return checklist, nil
}
