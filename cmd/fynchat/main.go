package main

import (
	"log"
	//	"fynchat/internal/api"
	"fynchat/internal/handlers"
	"fynchat/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	pingHandler := handlers.NewPingHandler()

	r := gin.Default()

	routes.Setup(r, pingHandler)

	if err := r.Run(":" + "8080"); err != nil {
		log.Fatal(err)
	}
}
