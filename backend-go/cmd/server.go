package main

import (
	"backend/config"
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/routes"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.SetupDB()
	r := gin.Default()


	repository := repository.NewUserRepository(db)
	service := service.NewUserService(db, repository)
	handler := handlers.NewUserHandler(service)

	routes.SetupRouter(r, &routes.Routes{
		UserHandler: handler,
	})

	r.Run(":8080")
}