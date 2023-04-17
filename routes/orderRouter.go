package routes
import(
	"github.com/gin-gonic/gin"
	controller "restaurant/controllers"
)
func OrderRoutes( engine *gin.Engine){
	engine.GET("/orders" ,controller.GetOrders())
	engine.GET("/orders/:order_id",controller.GetOrder())
	engine.POST("/orders",controller.CreateOrder())
	engine.PATCH("/orders/:order_id",controller.UpdateOrder())
}