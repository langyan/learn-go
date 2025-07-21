package repository

import (
	"context"
	"database/sql"

	"github.com/langyan/lin-go-service-layer/internal/model"
)

type UserRepository interface {
	Save(ctx context.Context, user *model.User) error
}

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) Save(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id, email) VALUES ($1,$2)", user.ID,  user.Email)
	return err
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}
