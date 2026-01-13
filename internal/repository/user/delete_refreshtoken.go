package user

import (
	"context"
	"errors"
)

func (r *userRepository) DeleteRefreshToken(ctx context.Context, userID int64) error {
	// delete the refresh token by user id

	query := `DELETE FROM refresh_tokens
	WHERE user_id=?`

	result, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("nothing to delete")
	}
	if rowsAffected == 0 {
		return errors.New("nothing to delete")
	}
	return nil
}
