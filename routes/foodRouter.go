package routes

import (
	"restaurant_management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/food", controllers.GetFoods())
	incomingRoutes.GET("/food/:food_id", controllers.GetFood())
	incomingRoutes.POST("/food", controllers.CreateFood())
	incomingRoutes.PATCH("/food/:food_id", controllers.UpdateFood())
}
