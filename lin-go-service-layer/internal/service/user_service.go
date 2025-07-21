package service

import (
	"context"

	"errors"

	"github.com/google/uuid"
	"github.com/langyan/lin-go-service-layer/internal/model"
	"github.com/langyan/lin-go-service-layer/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req model.CreateUserRequest) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req model.CreateUserRequest) (*model.User, error) {
	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	user := &model.User{
		ID:    uuid.New().String(),
		Email: req.Email,
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
