package post

import (
	"context"
	"errors"
	"net/http"
	"time"
	"tweets/internal/models"
)

func (s *postService) LikeOrUnlikePost(ctx context.Context, postID, userID int64) (int, error) {
	// check post was exists
	postExist, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if postExist == nil {
		return http.StatusNotFound, errors.New("tweet not found")
	}

	// check user already like or not
	isUserAlreadyLikePost, err := s.postRepo.IsUserAlreadyLikePost(ctx, postID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// if user already like, delete data
	if isUserAlreadyLikePost {
		err := s.postRepo.DeleteLikePost(ctx, postID, userID)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		// else, store data
		now := time.Now()
		err := s.postRepo.StoreLikePost(ctx, &models.PostLikeModel{
			UserID:    userID,
			PostID:    postID,
			CreatedAt: now,
			UpdatedAt: now,
		})
		if err != nil {
			return http.StatusInternalServerError, err
		}

	}

	// return
	return http.StatusOK, nil
}
