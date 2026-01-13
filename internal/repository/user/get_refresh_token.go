package user

import (
	"context"
	"database/sql"
	"time"
	"tweets/internal/models"
)

func (r *userRepository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*models.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, expired_at
	FROM refresh_tokens
	WHERE user_id = ? AND expired_at >= ?`

	row := r.db.QueryRowContext(ctx, query, userID, now)
	var result models.RefreshTokenModel
	err := row.Scan(&result.ID, &result.UserID, &result.RefreshToken, &result.ExpiredAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
