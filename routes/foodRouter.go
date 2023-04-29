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
func FavoriteRoutes(engine *gin.Engine){
	engine.GET("/favorites/:user_id" ,controller.GetFavorites())
 	engine.POST("/favorites",controller.CreateFavorite())
	engine.DELETE("/favorites/:favorite_id",controller.DeleteFavorite())
}