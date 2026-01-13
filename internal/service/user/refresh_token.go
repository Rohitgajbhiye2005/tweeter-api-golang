package user

import (
	"context"
	"errors"
	"net/http"
	"time"
	"tweets/internal/dto"
	"tweets/internal/models"
	"tweets/pkg/jwt"
	"tweets/pkg/refreshtoken"
)




func (s *userService)RefreshToken(ctx context.Context,req *dto.RefreshTokenRequest,userID int64)(string,string,int,error){
	// check user exits
	userExists,err:=s.userRepo.GetUserByID(ctx,userID)

	if err!=nil{
		return "","",http.StatusInternalServerError,err
	}

	if userExists==nil{
		return "","",http.StatusNotFound,errors.New("User not found")
	}


	// get refresh token by user id
	refreshTokenExists,err:=s.userRepo.GetRefreshToken(ctx,userID,time.Now())
	if err!=nil{
		return "","",http.StatusInternalServerError,errors.New("error in geting refresh token")
	}
	if refreshTokenExists==nil{
		return "","",http.StatusNotFound,errors.New("Refresh token was expired")	
	}

	// check refresh token is matched with request body

	if req.RefreshToken!=refreshTokenExists.RefreshToken{
		return 	 "","",http.StatusNotFound,errors.New("Refresh token not found")
	}

	// generate new token

	token,err:=jwt.CreateToken(userID,userExists.Username,s.cfg.SecretJwt)
	if err!=nil{
		return "","",http.StatusInternalServerError,err
	}

	// delete old refresh token & generate new refresh token

	err=s.userRepo.DeleteRefreshToken(ctx,userID)
	if err!=nil{
		return "","",http.StatusInternalServerError,err
	}

	refreshToken,err:=refreshtoken.GenerateRefreshToken()
	if err!=nil{
		return "","",http.StatusInternalServerError,err
	}

	now:=time.Now()
	s.userRepo.StoreRefreshToken(ctx,&models.RefreshTokenModel{
		UserID: userID,
		RefreshToken: refreshToken,
		CreatedAt: now,
		UpdatedAt: now,
		ExpiredAt: time.Now().Add(7*24*time.Hour),
	})

	return token,refreshToken,http.StatusOK,nil



}