package main

import (
	"fmt"
	"log"
	"net/http"

	"fynchat/internal/api"
	"fynchat/internal/database"
)

func main() {
	db := database.GetConnection()
	defer db.Close()

	routesHandler := api.RegisterRoutes(db)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8085", routesHandler))
}
