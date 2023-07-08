package routes

import (
	"alura-backend-7/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	depositions := r.Group("/depoimentos")
	{
		depositions.POST("/", controllers.CreateDeposition)
		depositions.GET("/", controllers.GetAllDepositions)
		depositions.GET("/:id", controllers.GetDepositionById)
		depositions.PUT("/:id", controllers.UpdateDeposition)
		depositions.DELETE("/:id", controllers.DeleteDeposition)
	}
}
