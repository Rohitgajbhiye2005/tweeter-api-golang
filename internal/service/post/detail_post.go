package post

import (
	"context"
	"errors"
	"net/http"
	"tweets/internal/dto"
)

func (s *postService)DetailPost(ctx context.Context,postID int64)(*dto.DetailPostResponse,int,error){
	// get post by id
	post,err:=s.postRepo.GetPostByID(ctx,postID)
	if err!=nil{
		return nil,http.StatusInternalServerError,err
	}

	if post == nil{
		return nil,http.StatusNotFound,errors.New("tweet not found")
	}

	// get all comment related to the id

	postIDs:=[]int64{post.ID}
	comments,err:=s.commentRepo.GetCommentsByPostIDs(ctx,postIDs)
	if err!=nil{
		return nil,http.StatusInternalServerError,err
	}


	// mapping comments with post

	commentsMap:=make([]dto.Comment,0)

	for _,comment:=range comments{
		commentsMap=append(commentsMap, dto.Comment{
			ID: comment.ID,
			Username: comment.Username,
			Content: comment.Content,
			LikeCount: comment.LikeCount,
			CreatedAt: comment.CreatedAt.String(),
			UpdatedAt: comment.UpdatedAt.String(),
		})
	}

	// set reponse

	return &dto.DetailPostResponse{
			ID: post.ID,
			Username: post.Username,
			Title: post.Title,
			Content: post.Content,
			LikeCount: post.LikeCount,
			Comments: commentsMap,
			CreatedAt: post.CreatedAt.String(),
			UpdatedAt: post.UpdatedAt.String(),
	},http.StatusOK,nil
}