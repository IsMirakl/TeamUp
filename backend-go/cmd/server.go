package main

import (
	postroutes "backend/internal/features/post"
	createpost "backend/internal/features/post/create_post"
	getauthorpost "backend/internal/features/post/get_author_post"
	getpostbyid "backend/internal/features/post/get_by_id"
	updatepost "backend/internal/features/post/update_post"
	userroutes "backend/internal/features/user"
	getuserbyemail "backend/internal/features/user/get_by_email"
	getuserbyid "backend/internal/features/user/get_by_id"
	loginuser "backend/internal/features/user/login_user"
	registeruser "backend/internal/features/user/register_user"
	"backend/internal/pkg/config"
	"time"

	"backend/internal/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	log := logger.NewLogger()
	cfg := config.New(log)
	signingKey := []byte(cfg.SECRET_KEY.JWT_SECRET)

	db := config.SetupDB()
	defer db.Pool.Close()
	r := gin.Default()


	log.Info("Server started")

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"PUT", "GET", "POST", "PATCH",
			"DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	registerRepo := registeruser.NewUserRepository(db.Queries, db.Pool)
	registerService := registeruser.NewUserService(registerRepo, log)
	registerHandler := registeruser.NewUserHandler(registerService, log)


	getUserByIdRepo := getuserbyid.NewRepository(db.Queries)
	getUserByIdService := getuserbyid.NewService(getUserByIdRepo, log)
	getUserByIdHandler := getuserbyid.NewHandler(getUserByIdService, log)

	getUserByEmailRepo := getuserbyemail.NewRepository(db.Queries)
	getUserByEmailService := getuserbyemail.NewService(getUserByEmailRepo, log)
	getUserByEmailHandler := getuserbyemail.NewHandler(getUserByEmailService, log)

	loginRepo := loginuser.NewRepository(db.Queries)
	loginService := loginuser.NewUserService(loginRepo, log)
	loginHandler := loginuser.NewUserHandler(loginService, log)

	createPostRepo := createpost.NewRepository(db.Queries, db.Pool)
	createPostService := createpost.NewService(createPostRepo, log)
	createPostHandler := createpost.NewHandler(createPostService, log)

	updatePostRepo := updatepost.NewRepository(db.Queries, db.Pool)
	updatePostService := updatepost.NewService(updatePostRepo, log)
	updatePostHandler := updatepost.NewHandler(updatePostService)

	getPostByIdRepo := getpostbyid.NewRepository(db.Queries)
	getPostByIdService := getpostbyid.NewService(getPostByIdRepo, log)
	getPostByIdHandler := getpostbyid.NewHandler(getPostByIdService, log)

	getAuthorPostRepo := getauthorpost.NewRepository(db.Queries)
	getAuthorPostService := getauthorpost.NewService(getAuthorPostRepo, log)
	getAuthorPostHandler := getauthorpost.NewHandler(getAuthorPostService, log)

	api := r.Group("/api")
	userroutes.UserRouter(api, registerHandler, loginHandler, getUserByIdHandler, getUserByEmailHandler)
	postroutes.PostRouter(
		api,
		createPostHandler,
		updatePostHandler,
		getPostByIdHandler,
		getAuthorPostHandler,
		signingKey,
		log,
	)

	r.Run(":8080")
}
