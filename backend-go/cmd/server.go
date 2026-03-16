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

	repository := repository.NewUserRepository(db)
	service := service.NewUserService(db, repository)
	handler := handlers.NewUserHandler(service)

	routes.SetupRouter(r, &routes.Routes{
		UserHandler: handler,
	})

	r.Run(":8080")
}