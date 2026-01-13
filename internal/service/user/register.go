package user

import (
	"context"
	"errors"
	"net/http"
	"time"
	"tweets/internal/dto"
	"tweets/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error) {
	//check the user is already exits
	userExits, err := s.userRepo.GetUserByEmailOrUsername(ctx, req.Email, req.Username)

	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	if userExits != nil {
		return 0, http.StatusBadRequest, errors.New("User already exist")
	}

	//hash password

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, http.StatusInternalServerError, errors.New("Error in hashing the Password")
	}

	//create user
	userModel := &models.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(passwordHash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userID, err := s.userRepo.CreateUser(ctx, userModel)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	return userID, http.StatusCreated, nil
}
