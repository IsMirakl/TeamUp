package main

import (
	"time"

	postcreateapp "backend/internal/content/application/command/create_post"
	postupdateapp "backend/internal/content/application/command/update_post"
	postgetauthorapp "backend/internal/content/application/query/get_author_post"
	postgetbyidapp "backend/internal/content/application/query/get_by_id"
	postcreateinfra "backend/internal/content/infrastructure/persistence/create_post"
	postgetauthorinfra "backend/internal/content/infrastructure/persistence/get_author_post"
	postgetbyidinfra "backend/internal/content/infrastructure/persistence/get_by_id"
	postupdateinfra "backend/internal/content/infrastructure/persistence/update_post"
	postroutes "backend/internal/content/interfaces/http"
	createpost "backend/internal/content/interfaces/http/create_post"
	getauthorpost "backend/internal/content/interfaces/http/get_author_post"
	getpostbyid "backend/internal/content/interfaces/http/get_by_id"
	updatepost "backend/internal/content/interfaces/http/update_post"
	sessionapp "backend/internal/identity/application/command/create_session"
	userloginapp "backend/internal/identity/application/command/login_user"
	refreshsessionapp "backend/internal/identity/application/command/refresh_session"
	userregisterapp "backend/internal/identity/application/command/register_user"
	usergetbyemailapp "backend/internal/identity/application/query/get_by_email"
	usergetbyidapp "backend/internal/identity/application/query/get_by_id"
	getmyprofileapp "backend/internal/identity/application/query/get_my_profile"
	sessioninfra "backend/internal/identity/infrastructure/persistence/create_session"
	usergetbyemailinfra "backend/internal/identity/infrastructure/persistence/get_by_email"
	usergetbyidinfra "backend/internal/identity/infrastructure/persistence/get_by_id"
	getmyprofileinfra "backend/internal/identity/infrastructure/persistence/get_my_profile"
	userlogininfra "backend/internal/identity/infrastructure/persistence/login_user"
	refreshsessioninfra "backend/internal/identity/infrastructure/persistence/refresh_session"
	userregisterinfra "backend/internal/identity/infrastructure/persistence/register_user"
	userroutes "backend/internal/identity/interfaces/http"
	getuserbyemail "backend/internal/identity/interfaces/http/get_by_email"
	getuserbyid "backend/internal/identity/interfaces/http/get_by_id"
	getmyprofile "backend/internal/identity/interfaces/http/get_my_profile"
	loginuser "backend/internal/identity/interfaces/http/login_user"
	refreshsession "backend/internal/identity/interfaces/http/refresh_session"
	registeruser "backend/internal/identity/interfaces/http/register_user"

	"backend/internal/pkg/config"
	"backend/internal/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	log := logger.NewLogger()
	cfg, err := config.New(log)
	if err != nil {
		log.Fatal(err)
	}

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

	getUserByIdRepo := usergetbyidinfra.NewRepository(db.Queries)
	getUserByIdService := usergetbyidapp.NewService(getUserByIdRepo, log)
	getUserByIdHandler := getuserbyid.NewHandler(getUserByIdService, log)

	getUserByEmailRepo := usergetbyemailinfra.NewRepository(db.Queries)
	getUserByEmailService := usergetbyemailapp.NewService(getUserByEmailRepo, log)
	getUserByEmailHandler := getuserbyemail.NewHandler(getUserByEmailService, log)

	sessionRepo := sessioninfra.NewRepository(db.Queries, log)
	sessionService := sessionapp.NewSesssionService(sessionRepo, log)

	registerRepo := userregisterinfra.NewUserRepository(db.Queries, db.Pool)
	registerService := userregisterapp.NewUserService(registerRepo, sessionService, log)
	registerHandler := registeruser.NewUserHandler(registerService, log)

	loginRepo := userlogininfra.NewRepository(db.Queries)
	loginService := userloginapp.NewUserService(loginRepo, sessionService, log)
	loginHandler := loginuser.NewUserHandler(loginService, log)

	refreshSessionRepo := refreshsessioninfra.NewRepository(db.Queries, log)
	refreshSessionService := refreshsessionapp.NewSessionService(refreshSessionRepo, log)
	refreshSessionHandler := refreshsession.NewHandler(refreshSessionService, log)

	createPostRepo := postcreateinfra.NewRepository(db.Queries, db.Pool)
	createPostService := postcreateapp.NewService(createPostRepo, log)
	createPostHandler := createpost.NewHandler(createPostService, log)

	updatePostRepo := postupdateinfra.NewRepository(db.Queries, db.Pool)
	updatePostService := postupdateapp.NewService(updatePostRepo, log)
	updatePostHandler := updatepost.NewHandler(updatePostService)

	getPostByIdRepo := postgetbyidinfra.NewRepository(db.Queries)
	getPostByIdService := postgetbyidapp.NewService(getPostByIdRepo, log)
	getPostByIdHandler := getpostbyid.NewHandler(getPostByIdService, log)

	getAuthorPostRepo := postgetauthorinfra.NewRepository(db.Queries)
	getAuthorPostService := postgetauthorapp.NewService(getAuthorPostRepo, log)
	getAuthorPostHandler := getauthorpost.NewHandler(getAuthorPostService, log)

	getProfileMeRepo := getmyprofileinfra.NewRepository(db.Queries, log)
	getProfileMeService := getmyprofileapp.NewPostService(getProfileMeRepo, log)
	getMyProfileHandler := getmyprofile.NewProfileHandler(getProfileMeService, log)

	api := r.Group("/api")
	userRouterParams := userroutes.NewRouterParams(signingKey, log)
	userroutes.UserRouter(api, registerHandler, loginHandler, refreshSessionHandler, getUserByIdHandler, getUserByEmailHandler, userRouterParams, getMyProfileHandler)
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
