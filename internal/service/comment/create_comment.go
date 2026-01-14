package comment

import (
	"context"
	"errors"
	"net/http"
	"time"
	"tweets/internal/dto"
	"tweets/internal/models"
)



func (s *commentService)CreateComment(ctx context.Context,req *dto.StoreCommentRequest,userID int64)(int,error){

	// check tweet is exists
	postExist,err:=s.postRepo.GetPostByID(ctx,req.PostID)
	if err!=nil{
		return http.StatusInternalServerError,err
	}

	if postExist==nil{
		return http.StatusNotFound,errors.New("tweet not found")
	}


	// store comment

	now:=time.Now()

	err=s.commentRepo.StoreComment(ctx,&models.CommentModel{
		PostID: req.PostID,
		UserID: userID,
		Content: req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err!=nil{
		return http.StatusInternalServerError,err
	}
	

	// return
	return http.StatusCreated,nil

}