package user

import (
	"context"
	"tweets/internal/models"
)


func (r *userRepository)StoreRefreshToken(ctx context.Context,models *models.RefreshTokenModel)error{
	query:=`INSERT INTO refresh_tokens (user_id, refresh_token, expired_at, created_at, updated_at)
	VALUES (?,?,?,?,?)`

	_,err:=r.db.ExecContext(ctx,query,models.UserID,models.RefreshToken,models.ExpiredAt,models.CreatedAt,models.UpdatedAt)

	return err
}