package routes
import(
	"github.com/gin-gonic/gin"
	controller "restaurant/controllers"
)
func InvoiceRoutes( engine *gin.Engine){
	engine.GET("/invoices" ,controller.GetInvoices())
	engine.GET("/invoices/:invoice_id",controller.GetInvoice())
	engine.POST("/invoices",controller.CreateInvoice())
	engine.PATCH("/invoices/:invoice_id",controller.UpdateInvoice())
}