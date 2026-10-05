package routes

import (
	"fynchat/internal/handlers"
	"fynchat/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine,
	pingHandler *handlers.PingHandler,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	refreshHandler *handlers.RefreshHandler,
	channelHandler *handlers.ChannelHandler,
	messagesHandler *handlers.MessageHandler,
	jwtSecret string) {

	r.GET("/ping", pingHandler.Ping)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/logout", refreshHandler.Logout)

	api := r.Group("/api")
	api.Use(middlewares.Auth(jwtSecret))
	{
		api.GET("/me", userHandler.GetMe)
		api.POST("/user/update", userHandler.UpdateMe)
		api.POST("/user/delete", userHandler.DeleteMe)
		api.POST("/request", refreshHandler.Refresh)
		api.POST("/channel/create", channelHandler.Create)
		api.POST("/message/create", messagesHandler.Create)
		// api.GET("/channel/:id", channelHandler.GetChannel)
	}
}
