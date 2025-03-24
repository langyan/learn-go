package services

import (
	"context"

	"github.com/langyan/web-gin-gorm-redis/models"
	"github.com/langyan/web-gin-gorm-redis/repositories"
)

func GetUser(ctx context.Context, id uint) (*models.User, error) {
	return repositories.GetUserByID(ctx, id)
}
