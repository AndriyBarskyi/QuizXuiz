// routes/routes.go
package routes

import (
	"QuizXuiz/controllers"
	"QuizXuiz/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(secret string) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.RegisterUser)
			auth.POST("/login", controllers.LoginUser)
		}
		secured := api.Group("/user", middleware.AuthMiddleware(secret))
		{
			secured.GET("/profile", controllers.GetUserProfile)
		}
	}
	return r
}
