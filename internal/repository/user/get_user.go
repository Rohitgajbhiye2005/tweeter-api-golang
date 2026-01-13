package user

import (
	"context"
	"database/sql"
	"tweets/internal/models"
)

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*models.UserModel, error) {
	query := `SELECT id,username,email,created_at,updated_at
	FROM users WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, userID)
	var result models.UserModel
	err := row.Scan(&result.ID, &result.Username, &result.Email, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
