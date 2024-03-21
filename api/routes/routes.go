package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router = createUserRoutes(router)
}

func createUserRoutes(router *gin.Engine) *gin.Engine {
	user_api := NewUser()
	router.POST("/api/users", user_api.CreateUserHandler)
	router.GET("/api/users/:id", user_api.GetUserById)
	router.PUT("/api/users/:id", user_api.ModifyUserById)
	router.DELETE("/api/users/:id", user_api.DeleteUserById)
	router.POST("/api/users/search", user_api.SearchUser)
	return router
}
