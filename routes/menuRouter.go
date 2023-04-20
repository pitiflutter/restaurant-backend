package routes
import(
	"github.com/gin-gonic/gin"
	controller "restaurant/controllers"
)
func MenuRoutes( engine *gin.Engine){
	engine.GET("/menus" ,controller.GetMenus())
	engine.GET("/menus/:menus_id",controller.GetMenu())
	engine.POST("/menus",controller.CreateMenu())
	engine.PATCH("menus/:menu_id",controller.UpdateMenu())
}