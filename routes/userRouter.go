package routes
import(
	"github.com/gin-gonic/gin"
	controller "restaurant/controllers"
)
func UserRoutes(engine *gin.Engine){
	engine.GET("/users" ,controller.GetUsers())
	engine.GET("/users/:user_id",controller.GetUser())
	
}
func AuthRoutes(engine *gin.Engine){
	engine.POST("/users/signup",controller.SignUp())
	engine.POST("/users/login",controller.Login())
}