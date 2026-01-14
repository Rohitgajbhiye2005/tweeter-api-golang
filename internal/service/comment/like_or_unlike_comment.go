package comment

import (
	"context"
	"errors"
	"net/http"
	"time"
	"tweets/internal/models"
)

func (s *commentService) LikeOrUnlikeComment(ctx context.Context, commentID, userID int64) (int, error) {
	// check comment is exists
	commentExist, err := s.commentRepo.DetailComment(ctx, commentID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if commentExist == nil {
		return http.StatusNotFound, errors.New("comment not found")
	}

	// check user already like the comment

	isUserAlreadyLikeComment, err := s.commentRepo.IsUserAlreadyLikeComment(ctx, commentID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// if user was like, delete data
	if isUserAlreadyLikeComment {
		err := s.commentRepo.DeleteLikeComment(ctx, commentID, userID)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	} else {
		// else, store data
		now := time.Now()
		err := s.commentRepo.StoreLikeComment(ctx, &models.CommentLikeModel{
			UserID:    userID,
			CommentID: commentID,
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
