package main

import (
	"fynchat.thegrigoriyagen.com/internal/config"
	"fynchat.thegrigoriyagen.com/internal/database"
	"fynchat.thegrigoriyagen.com/internal/handlers"
	"fynchat.thegrigoriyagen.com/internal/repository"
	"fynchat.thegrigoriyagen.com/internal/routes"
	"fynchat.thegrigoriyagen.com/internal/services"
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
	messagesRepo := repository.NewMessagesRepository(pool)

	userService := services.NewUserService(userRepo)
	refreshService := services.NewRefreshService(refreshRepo)
	channelService := services.NewChannelService(channelRepo)
	messageService := services.NewMessagesService(messagesRepo)

	pingHandler := handlers.NewPingHandler()
	authHandler := handlers.NewAuthHandler(userService, cfg.JWTSecret)
	userHandler := handlers.NewUserHandler(userService)
	refreshHandler := handlers.NewRefreshHandler(refreshService, cfg.JWTSecret)
	channelHandler := handlers.NewChannelHandler(channelService)
	messagesHandler := handlers.NewMessagesHandler(messageService)

	r := gin.Default()

	routes.Setup(r, pingHandler, authHandler, userHandler, refreshHandler, channelHandler, messagesHandler, cfg.JWTSecret)

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
