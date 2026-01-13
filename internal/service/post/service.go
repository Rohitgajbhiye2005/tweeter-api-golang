package post

import (
	"context"
	"tweets/internal/config"
	"tweets/internal/dto"
	"tweets/internal/repository/post"
)


type PostService interface{
	CreatePost(ctx context.Context,req *dto.CreatePostRequest,userID int64)(int64,int ,error)
}

type postService struct{
	cfg *config.Config
	postRepo post.PostRepository
}

func NewPostService(cfg *config.Config,postRepo post.PostRepository) PostService {
	return &postService{
		cfg:cfg,
		postRepo:postRepo,
	}
}