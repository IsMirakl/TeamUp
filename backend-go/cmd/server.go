package main

import (
	"backend/config"
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/routes"
	"backend/internal/service"
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
		AllowHeaders: []string{"Origin", "Content-type", "Authorization", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

userRepo := repository.NewUserRepository(db)
userService := service.NewUserService(db, userRepo)
userHandler := handlers.NewUserHandler(userService)

teamSeekPostRepo := repository.NewTeamSeekPostRepository(db)
teamSeekPostService := service.NewTeamSeekPostService(db, teamSeekPostRepo)
teamSeekPostHandler := handlers.NewTeamSeekPostHandler(teamSeekPostService)

routes.SetupRouter(r, &routes.Routes{
	UserHandler:         userHandler,
	TeamSeekPostHandler: teamSeekPostHandler,
	SigningKey: signingKey,
})


	r.Run(":8080")
}