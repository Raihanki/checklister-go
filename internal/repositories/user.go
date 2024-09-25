package repositories

import (
	"database/sql"

	"github.com/Raihanki/checklisters/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	Create(ctx *fiber.Ctx, tx *sql.Tx, user domain.User) error
	GetByEmail(ctx *fiber.Ctx, db *sql.DB, email string) (domain.User, error)
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Create(ctx *fiber.Ctx, tx *sql.Tx, user domain.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := tx.ExecContext(ctx.Context(), query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (repo *UserRepositoryImpl) GetByEmail(ctx *fiber.Ctx, db *sql.DB, email string) (user domain.User, err error) {
	query := "SELECT id, name, email, password FROM users WHERE email = $1"
	err = db.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
