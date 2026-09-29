package main

import (
	"fmt"
	"gohst_api_example/middleware"
	"gohst_api_example/shared"
	"gohst_api_example/user"

	"github.com/cccaaannn/gohst"
)

func main() {
	// Create a new server
	server := gohst.CreateServer()

	// Set default headers
	headers := map[string]string{"Content-Type": "application/json"}
	server.SetHeaders(headers)

	// Add middlewares
	server.Use(middleware.Authentication)
	server.Use(middleware.Authorization)

	// Add user routes
	user.CreateUserRoutes(server)

	// Add a catch-all route for 404 Not Found
	shared.CreateNotFoundRoute(server)

	// Start the server
	stop, err := server.ListenAndServe(":8080")
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
	defer close(stop)
	<-stop
}
