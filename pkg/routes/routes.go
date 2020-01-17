package routes

import "github.com/gin-gonic/gin"

import "sample-login-service/pkg/createuser"

// CreateRoutes acts as a map of the service
func CreateRoutes(server *gin.Engine) {

	server.POST("/create/user", createuser.CreateUser)

}
