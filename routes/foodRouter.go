package routes
import(
	"github.com/gin-gonic/gin"
	controller "restaurant/controllers"
)
func FoodRoutes( engine *gin.Engine){
	engine.GET("/foods" ,controller.GetFoods())
	engine.GET("/foods/:food_id",controller.GetFood())
	engine.POST("/foods",controller.CreateFood())
	engine.PATCH("/foods/:food_id",controller.UpdateFood())
}