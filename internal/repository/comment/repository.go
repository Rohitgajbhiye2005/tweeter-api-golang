package comment

import (
	"context"
	"database/sql"
	"tweets/internal/models"
)


type CommentRepository interface{
	StoreComment(ctx context.Context,model *models.CommentModel)error
	DetailComment(ctx context.Context,commentID int64)(*models.CommentModel,error)
	IsUserAlreadyLikeComment(ctx context.Context,commentID,userID int64)(bool,error)
	DeleteLikeComment(ctx context.Context,commentID,userID int64)error
	StoreLikeComment(ctx context.Context,model *models.CommentLikeModel) error
	GetCommentsByPostIDs(ctx context.Context,postIDs []int64)([]models.CommentModel,error)

}

type commentRepository struct{
	db *sql.DB
}

func NewCommentRepository(db *sql.DB)CommentRepository{
	return &commentRepository{
		db: db,
	}
}