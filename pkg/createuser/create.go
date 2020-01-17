package createuser

import "github.com/gin-gonic/gin"

import "net/http"

// CreateUser will make take a username
// and a password and create a hash of the
// of the password using bcrypt
func CreateUser(c *gin.Context) {

	c.JSON(http.StatusOK, "This Seems to be working")
}
