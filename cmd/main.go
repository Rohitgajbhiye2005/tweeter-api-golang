package main

import (
	"fmt"
	"log"
	"tweets/internal/config"
	userHandler "tweets/internal/handlers/user"
	userRepo "tweets/internal/repository/user"
	postRepo "tweets/internal/repository/post"
	userService "tweets/internal/service/user"
	postHandler "tweets/internal/handlers/post"
	postService "tweets/internal/service/post"
	"tweets/pkg/internalsql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()
	validate:=validator.New()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := internalsql.ConnectMySQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "working",
		})
	})

	userRepo := userRepo.NewRepository(db)
	postRepo:=postRepo.NewPostRepository(db)

	userService := userService.NewService(cfg, userRepo)
	postService:=postService.NewPostService(cfg,postRepo)

	userHandler := userHandler.NewHandler(r, validate,userService)
	postHandler:=postHandler.NewHandler(r,validate,postService)

	userHandler.RouteList(cfg.SecretJwt)
	postHandler.RouteList(cfg.SecretJwt)

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)

	r.Run(server)
}
