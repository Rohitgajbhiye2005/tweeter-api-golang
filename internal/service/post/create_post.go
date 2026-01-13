package post

import (
	"context"
	"net/http"
	"time"
	"tweets/internal/dto"
	"tweets/internal/models"
)


func (s *postService)CreatePost(ctx context.Context,req *dto.CreatePostRequest,userID int64)(int64,int ,error){
	// store post
	now:=time.Now()
	insertedID,err:=s.postRepo.StorePost(ctx,&models.PostModel{
		UserID: userID,
		Title: req.Title,
		Content: req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err!=nil{
		return 0,http.StatusInternalServerError,err
	}
	// return

	return insertedID,http.StatusCreated,nil
}