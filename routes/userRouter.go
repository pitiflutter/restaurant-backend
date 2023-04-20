package routes
import(
	"github.com/gin-gonic/gin"
	controller "restaurant/controllers"
)
func UserRoutes(engine *gin.Engine){
	engine.GET("/users" ,controller.GetUsers())
	engine.GET("/users/:user_id",controller.GetUser())
	engine.POST("/users/signup",controller.SignUp())
	engine.POST("/users/login",controller.Login())
}