package main

import (
	postroutes "backend/internal/features/post"
	createpost "backend/internal/features/post/create_post"
	getauthorpost "backend/internal/features/post/get_author_post"
	getpostbyid "backend/internal/features/post/get_by_id"
	updatepost "backend/internal/features/post/update_post"
	userroutes "backend/internal/features/user"
	getuserbyemail "backend/internal/features/user/get_by_email"
	loginuser "backend/internal/features/user/login_user"
	registeruser "backend/internal/features/user/register_user"
	"backend/internal/pkg/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.New()
	signingKey := []byte(cfg.SECRET_KEY.JWT_SECRET)

	db := config.SetupDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"PUT", "GET", "POST", "PATCH",
			"DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	registerRepo := registeruser.NewUserRepository(db)
	registerService := registeruser.NewUserService(db, registerRepo)
	registerHandler := registeruser.NewUserHandler(registerService)

	getUserByEmailRepo := getuserbyemail.NewRepository(db)
	getUserByEmailService := getuserbyemail.NewService(getUserByEmailRepo)
	loginService := loginuser.NewUserService(getUserByEmailService)
	loginHandler := loginuser.NewUserHandler(loginService)

	createPostRepo := createpost.NewRepository(db)
	createPostService := createpost.NewService(db, createPostRepo)
	createPostHandler := createpost.NewHandler(createPostService)

	updatePostRepo := updatepost.NewRepository(db)
	updatePostService := updatepost.NewService(db, updatePostRepo)
	updatePostHandler := updatepost.NewHandler(updatePostService)

	getPostByIdRepo := getpostbyid.NewRepository(db)
	getPostByIdService := getpostbyid.NewService(getPostByIdRepo)
	getPostByIdHandler := getpostbyid.NewHandler(getPostByIdService)

	getAuthorPostRepo := getauthorpost.NewRepository(db)
	getAuthorPostService := getauthorpost.NewService(getAuthorPostRepo)
	getAuthorPostHandler := getauthorpost.NewHandler(getAuthorPostService)

	api := r.Group("/api")
	userroutes.UserRouter(api, registerHandler, loginHandler)
	postroutes.PostRouter(
		api,
		createPostHandler,
		updatePostHandler,
		getPostByIdHandler,
		getAuthorPostHandler,
		signingKey,
	)

	r.Run(":8080")
}
