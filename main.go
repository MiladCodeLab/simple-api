package main

import (
	"github.com/MiladCodeLab/simple-api/application"
	"github.com/MiladCodeLab/simple-api/repository"
	"github.com/MiladCodeLab/simple-api/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func main() {
	logger := slog.Default()

	repo := repository.NewUserRepository(logger)
	userService := service.NewUserService(logger, repo)
	userHandler := application.NewUserHandler(logger, userService)

	httpServer := gin.Default()
	userHandler.RegisterRoutes(httpServer)

	if err := httpServer.Run(":5001"); err != nil {
		logger.Error("error", err)
		os.Exit(1)
	}
}
