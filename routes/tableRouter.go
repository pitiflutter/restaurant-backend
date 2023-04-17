package routes
import(
	"github.com/gin-gonic/gin"
	controller "restaurant/controllers"
)
func TableRoutes( engine *gin.Engine){
	engine.GET("/tables" ,controller.GetTables())
	engine.GET("/tables/:table_id",controller.GetTable())
	engine.POST("/tables",controller.CreateTable())
	engine.PATCH("/tables/:table_id",controller.UpdateTable())
}