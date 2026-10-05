package main

import (
	"fynchat/internal/config"
	"fynchat/internal/database"
	"fynchat/internal/handlers"
	"fynchat/internal/repository"
	"fynchat/internal/routes"
	"fynchat/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	userRepo := repository.NewUserRepository(pool)
	refreshRepo := repository.NewRefreshRepository(pool)
	channelRepo := repository.NewChannelRepository(pool)

	userService := services.NewUserService(userRepo)
	refreshService := services.NewRefreshService(refreshRepo)
	channelService := services.NewChannelService(channelRepo)

	pingHandler := handlers.NewPingHandler()
	authHandler := handlers.NewAuthHandler(userService, cfg.JWTSecret)
	userHandler := handlers.NewUserHandler(userService)
	refreshHandler := handlers.NewRefreshHandler(refreshService, cfg.JWTSecret)
	channelHandler := handlers.NewChannelHandler(channelService)

	r := gin.Default()

	routes.Setup(r, pingHandler, authHandler, userHandler, refreshHandler, channelHandler, cfg.JWTSecret)

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
