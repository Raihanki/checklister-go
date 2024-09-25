package services

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Raihanki/checklisters/internal/domain"
	"github.com/Raihanki/checklisters/internal/dto"
	"github.com/Raihanki/checklisters/internal/errpkg"
	"github.com/Raihanki/checklisters/internal/helpers"
	"github.com/Raihanki/checklisters/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx *fiber.Ctx, request dto.LoginRequest) (dto.AuthResponse, error)
	Register(ctx *fiber.Ctx, request dto.RegisterRequest) (dto.AuthResponse, error)
}

type UserServiceImpl struct {
	DB             *sql.DB
	UserRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{UserRepository: userRepo, DB: db}
}

func (service *UserServiceImpl) Login(ctx *fiber.Ctx, request dto.LoginRequest) (dto.AuthResponse, error) {
	user, err := service.UserRepository.GetByEmail(ctx, service.DB, request.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.AuthResponse{}, errpkg.ErrInvalidEmailOrPassword
		}
		log.Printf("error while get user by email ERR::%v", err)
		return dto.AuthResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return dto.AuthResponse{}, errpkg.ErrInvalidEmailOrPassword
	}

	token, err := helpers.GenerateToken(user.Id)
	if err != nil {
		log.Printf("error while generating token ERR::%v", err)
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{
		Token:     token,
		ExpiredAt: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}, nil
}

func (service *UserServiceImpl) Register(ctx *fiber.Ctx, request dto.RegisterRequest) (dto.AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.AuthResponse{}, nil
	}
	user := domain.User{
		Email:    request.Email,
		Password: string(hashedPassword),
		Name:     request.Name,
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return dto.AuthResponse{}, nil
	}
	defer tx.Rollback()

	err = service.UserRepository.Create(ctx, tx, user)
	if err != nil {
		log.Printf("error while create user ERR::%v", err)
		return dto.AuthResponse{}, nil
	}

	token, err := helpers.GenerateToken(user.Id)
	if err != nil {
		log.Printf("error while generating token ERR::%v", err)
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{
		Token:     token,
		ExpiredAt: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}, nil
}
