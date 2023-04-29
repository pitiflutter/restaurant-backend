package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"restaurant/routes"
 
	"restaurant/middleware"
	 
	 
)
 

func main(){
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router :=gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
 	router.Use(middleware.Authentication())
	routes.UserRoutes(router)
 	routes.FoodRoutes(router)
	routes.FavoriteRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)


	router.Run()

}