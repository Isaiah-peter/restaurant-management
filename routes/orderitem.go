package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(route *gin.Engine) {
	route.GET("/orderitems", controllers.GetOrderItems())
	route.GET("/orderitem/:orderitem_id", controllers.GetOrderItem())
	route.POST("/orderitems", controllers.CreateOrderItem())
	route.POST("/orderitem/:orderitem_id", controllers.UpdateOrderItem())
}
