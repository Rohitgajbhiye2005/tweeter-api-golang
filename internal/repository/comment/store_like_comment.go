package comment

import (
	"context"
	"tweets/internal/models"
)

func (r *commentRepository) StoreLikeComment(ctx context.Context, model *models.CommentLikeModel) error {
	query := `INSERT INTO comment_likes (comment_id, user_id, created_at, updated_at)
	VALUES (?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.CommentID, model.UserID, model.CreatedAt, model.UpdatedAt)

	return err
}
