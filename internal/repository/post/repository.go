package post

import (
	"context"
	"database/sql"
	"tweets/internal/models"
)


type PostRepository interface{


	StorePost(ctx context.Context,models *models.PostModel)(int64,error)

}

type postRepository struct{
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db:db,
	}
}