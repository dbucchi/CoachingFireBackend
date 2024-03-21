package main

import (
	"CoachingFireBackend/api/routes"
	"CoachingFireBackend/internal/utility"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	utility.Configuration = *utility.NewConfiguration()
	utility.ApplicationCache = *utility.NewCache()
	router := gin.Default()

	routes.SetupRoutes(router)

	port := 8080
	fmt.Printf("Start server on port %d...\n", port)
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Failed to start server err:", err)
	}
}
