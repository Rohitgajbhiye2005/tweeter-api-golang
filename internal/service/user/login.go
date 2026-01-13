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

	"golang.org/x/crypto/bcrypt"
)

func (s *userService)Login(ctx context.Context,req *dto.LoginRequest)(string,string,int,error){
	// check if user is register or not
	userExits,err:=s.userRepo.GetUserByEmailOrUsername(ctx,req.Email,"")
	if err!=nil{
		return "","",http.StatusInternalServerError,err
	}
	if userExits==nil{
		return "","",http.StatusNotFound,err
	}

	err=bcrypt.CompareHashAndPassword([]byte(userExits.Password),[]byte(req.Password))
	if err!=nil{
		return "","",http.StatusNotFound,errors.New("Wrong email or password") 
	}
	// generate the token
	token,err:=jwt.CreateToken(userExits.ID,userExits.Username,s.cfg.SecretJwt)
	if err!=nil{
		return "","",http.StatusInternalServerError,err
	}

	// get refresh token if exist
	now:=time.Now()
	refreshTokenExist,err:=s.userRepo.GetRefreshToken(ctx,userExits.ID,now)
	if err!=nil{
		return "","",http.StatusInternalServerError,err
	}

	if refreshTokenExist!=nil{
		return token,refreshTokenExist.RefreshToken,http.StatusOK,nil
	}

	// generate the refresh token and save it to database

	refreshToken,err:=refreshtoken.GenerateRefreshToken()
	if err!=nil{
		return "","",http.StatusInternalServerError,err
	}

	err=s.userRepo.StoreRefreshToken(ctx,&models.RefreshTokenModel{
		UserID: userExits.ID,
		RefreshToken: refreshToken,
		CreatedAt: now,
		UpdatedAt: now,
		ExpiredAt: time.Now().Add(7*24*time.Hour),
	})
	if err!=nil{
		return "","",http.StatusInternalServerError,err
	}
	// return
	return token,refreshToken,http.StatusOK,nil
}
