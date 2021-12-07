package routes

import (
	"github.com/dogerescat/go-app/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := router()
	router.Run(":8080")
}

func router() *gin.Engine {
	router := gin.Default()
	users := router.Group("/api/v1/user")
	uc := controllers.UserController{}
	{
		users.GET("/:id", uc.Read)
		users.POST("/", uc.Create)
	}
	return router
}
