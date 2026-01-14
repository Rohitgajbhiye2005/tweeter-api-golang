package post

import (
	"context"
	"tweets/internal/dto"
	"tweets/internal/models"
)

func (r *postRepository) GetAllPost(ctx context.Context, param *dto.GetAllPostRequest, offset int) ([]models.PostWithUserModel, error) {
	query := `SELECT
	p.id,p.title,p.content,p.user_id,p.created_at,p.updated_at,u.username,COUNT(p1.id) as like_count
	FROM posts as p
	JOIN users as u ON u.id-p.user_id
	LEFT JOIN post_likes as p1 ON p1.post_id = p.id
	WHERE p.deleted_at IS NULL
	GROUP BY p.id,p.title,p.content,p.user_id,p.created_at,p.updated_at,u.username
	ORDER BY created_at DESC
	LIMIT ?
	OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, param.Limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]models.PostWithUserModel, 0)

	for rows.Next() {
		var data models.PostWithUserModel
		err := rows.Scan(&data.ID, &data.Title, &data.Content, &data.UserID, &data.CreatedAt, &data.UpdatedAt, &data.Username, &data.LikeCount)
		if err != nil {
			return nil, err
		}
		result = append(result, models.PostWithUserModel{
			ID:        data.ID,
			Username:  data.Username,
			Title:     data.Title,
			Content:   data.Content,
			LikeCount: data.LikeCount,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		})
	}
	return result, nil
}
