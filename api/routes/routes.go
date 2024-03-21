package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes collega tutti gli endpoint al router
func SetupRoutes(router *gin.Engine) {

	user_api := NewUser()

	router.POST("/api/users", user_api.CreateUserHandler)
	// Aggiungi altri endpoint secondo le tue esigenze
}
