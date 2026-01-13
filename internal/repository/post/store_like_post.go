package post

import (
	"context"
	"tweets/internal/models"
)

func (r *postRepository) StoreLikePost(ctx context.Context, model *models.PostLikeModel) error {
	query := `INSERT INTO post_likes ( post_id, user_id, created_at, updated_at)
	VALUES (?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.CreatedAt, model.UpdatedAt)

	return err

}
