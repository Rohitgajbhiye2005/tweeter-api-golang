package post

import (
	"context"
	"database/sql"
	"time"
	"tweets/internal/dto"
	"tweets/internal/models"
)


type PostRepository interface{
	StorePost(ctx context.Context,models *models.PostModel)(int64,error)
	GetPostByID(ctx context.Context,postID int64)(*models.PostWithUserModel,error)
	UpdatePost(ctx context.Context,model *models.PostModel,postID int64) error
	SoftDeletePost(ctx context.Context,postID int64,now time.Time)error
	IsUserAlreadyLikePost(ctc context.Context,postID,userID int64)(bool,error)
	DeleteLikePost(ctx context.Context,postID,userID int64)error
	StoreLikePost(ctx context.Context,model *models.PostLikeModel)error
	TotalPost(ctx context.Context)(int64,error)
	GetAllPost(ctx context.Context,param *dto.GetAllPostRequest,offset int)([]models.PostWithUserModel,error)
}

type postRepository struct{
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db:db,
	}
}