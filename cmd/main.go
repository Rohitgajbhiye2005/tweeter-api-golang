package main

import (
	"fmt"
	"log"
	"tweets/internal/config"
	commentHandler "tweets/internal/handlers/comment"
	postHandler "tweets/internal/handlers/post"
	userHandler "tweets/internal/handlers/user"
	commentRep "tweets/internal/repository/comment"
	postRep "tweets/internal/repository/post"
	userRep "tweets/internal/repository/user"
	commentService "tweets/internal/service/comment"
	postService "tweets/internal/service/post"
	userService "tweets/internal/service/user"
	"tweets/pkg/internalsql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()
	validate := validator.New()
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

	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "working",
		})
	})

	userRepo := userRep.NewRepository(db)
	postRepo := postRep.NewPostRepository(db)
	commentRepo := commentRep.NewCommentRepository(db)

	userService := userService.NewService(cfg, userRepo)
	postService := postService.NewPostService(cfg, postRepo, commentRepo)
	commentService := commentService.NewCommentService(cfg, commentRepo, postRepo)

	userHandler := userHandler.NewHandler(r, validate, userService)
	postHandler := postHandler.NewHandler(r, validate, postService)
	commentHandler := commentHandler.NewHandler(r, validate, commentService)

	userHandler.RouteList(cfg.SecretJwt)
	postHandler.RouteList(cfg.SecretJwt)
	commentHandler.RouteList(cfg.SecretJwt)

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)

	r.Run(server)
}

// func main() {
// 	r := gin.Default()
// 	validate := validator.New()

// 	cfg, err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db, err := internalsql.ConnectMySQL(cfg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	r.Use(gin.Logger())
// 	r.Use(gin.Recovery())

// 	r.GET("/", func(ctx *gin.Context) {
// 		ctx.JSON(200, gin.H{
// 			"message": "working",
// 		})
// 	})

// 	// repositories (use different variable names)
// 	userRepository := userRepo.NewRepository(db)
// 	postRepository := postRepo.NewPostRepository(db)
// 	commentRepository := commentRepo.CommentRepository(db) // see note below

// 	// services
// 	userSvc := userService.NewService(cfg, userRepository)
// 	postSvc := postService.NewPostService(cfg, postRepository)
// 	commentSvc := commentService.NewCommentService(cfg, commentRepository)

// 	// handlers
// 	userH := userHandler.NewHandler(r, validate, userSvc)
// 	postH := postHandler.NewHandler(r, validate, postSvc)
// 	commentH:=

// 	userH.RouteList(cfg.SecretJwt)
// 	postH.RouteList(cfg.SecretJwt)

// 	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
// 	r.Run(server)
// }
