package main

import (
	"github.com/gin-gonic/gin"
	"jwt-auth-tejdeep/controllers"
	"jwt-auth-tejdeep/initializers"
	"jwt-auth-tejdeep/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "JWT Auth API by Tejdeep is running 🚀",
		})
	})

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	router.Run()
}
