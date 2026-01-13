package user

import (
	"context"
	"database/sql"
	"time"
	"tweets/internal/models"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email, username string) (*models.UserModel, error)
	CreateUser(ctx context.Context, model *models.UserModel) (int64, error)
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*models.RefreshTokenModel, error)
	StoreRefreshToken(ctx context.Context, models *models.RefreshTokenModel) error
	GetUserByID(ctx context.Context, userID int64) (*models.UserModel, error)
	DeleteRefreshToken(ctx context.Context, userID int64) error
}



type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
