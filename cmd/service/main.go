package main

import (
	"sample-login-service/pkg/routes"
	"sample-login-service/pkg/server"
)

func main() {

	// Make a new service
	s := server.CreateServer("sample-login-service")

	// Map the routes to handlers
	routes.CreateRoutes(s.Server)

	// Start the service
	s.Server.Run(":8080")
}
