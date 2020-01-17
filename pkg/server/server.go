package server

import (
	"github.com/gin-gonic/gin"
)

// Server is the structure that holds basic
// about the server name and  configurations
type Server struct {
	Server *gin.Engine
	Name   string
}

// CreateServer makes a new server and runs it
func CreateServer(name string) *Server {
	router := gin.New()

	router.Use(gin.Recovery())

	server := &Server{Server: router, Name: name}

	return server
}
