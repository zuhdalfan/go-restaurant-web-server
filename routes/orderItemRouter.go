package routes

import (
	"restaurant_management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItem", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItem/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItem-order/:order_id", controllers.GetOrderItemByOrder())
	incomingRoutes.POST("/orderItem", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItem/:orderItem_id", controllers.UpdateOrderItem())
}
