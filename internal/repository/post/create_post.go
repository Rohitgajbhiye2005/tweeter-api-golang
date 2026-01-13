package post

import (
	"context"
	"tweets/internal/models"
)

func (r *postRepository) StorePost(ctx context.Context, models *models.PostModel) (int64, error) {
	query := `INSERT INTO posts (user_id,title,content,created_at,updated_at)
	VALUES (?,?,?,?,?)`

	result, err := r.db.ExecContext(ctx, query, models.UserID, models.Title, models.Content, models.CreatedAt, models.UpdatedAt)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
