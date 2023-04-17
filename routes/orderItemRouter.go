package routes
import(
	"github.com/gin-gonic/gin"
	controller "restaurant/controllers"
)
func OrderItemRoutes( engine *gin.Engine){
	engine.GET("/orderItems" ,controller.GetOrderItems())
	engine.GET("/orderItems/:orderItem_id",controller.GetOrderItem())
	engine.GET("/orderItems-order/:order_id",controller.GetOrderItemsByOrder())
	engine.POST("/orderItemItems",controller.CreateOrderItem())
	engine.PATCH("/orderItems/:orderItem_id",controller.UpdateOrderItem())
}