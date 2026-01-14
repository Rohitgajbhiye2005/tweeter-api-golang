package comment

import (
	"context"
	"database/sql"
	"tweets/internal/models"
)

func (r *commentRepository)DetailComment(ctx context.Context,commentID int64)(*models.CommentModel,error){
	query:=`SELECT id, post_id, user_id, content, created_at, updated_at
	FROM comments
	WHERE id = ?`

	row:=r.db.QueryRowContext(ctx,query,commentID)
	var result models.CommentModel
	err:=row.Scan(&result.ID,&result.PostID,&result.UserID,&result.Content,&result.CreatedAt,&result.UpdatedAt)
	if err!=nil{
		if err==sql.ErrNoRows{
			return nil,nil
		}
		return nil,err
	}
	return &result,nil
}