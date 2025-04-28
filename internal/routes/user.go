// internal/routes/routes.go
package routes

import (
	middlewares "jwt-auth-server/internal/auth"
	"jwt-auth-server/internal/controllers"
	"jwt-auth-server/internal/repositories"
	"jwt-auth-server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Set up the router
	userRepo := repositories.NewUsertRepository()

	userService := services.NewUserService(userRepo)

	userHandler := controllers.NewUserHandler(userService)

	// check home
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Health check!"})
	})
	// User routes
	users := r.Group("/api/v1")
	{
		users.POST("/register", userHandler.CreateUser)
		users.GET("/users", userHandler.GetAllUsers)
		users.GET("/users/:id", userHandler.GetUserByID)
		users.PUT("/users/:id", userHandler.UpdateUser)
		users.DELETE("/users/:id", userHandler.DeleteUser)
	}
	// Auth routes
	auth := r.Group("/api/auth")
	{
		auth.POST("", userHandler.Login)
		auth.GET("/user", middlewares.JwtAuthMiddleware(), userHandler.CurrentUser)

	}

}
