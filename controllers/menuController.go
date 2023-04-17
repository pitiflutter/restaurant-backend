package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant/database"
	"restaurant/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/check.v1"
)
var menuCollection *mongo.Collection = database.OpenCollection(database.Client,"menu")

func GetMenus()gin.HandlerFunc{
	return func (c *gin.Context)  {
	ctx ,cancel := context.WithTimeout(context.Background(),100 *time.Second)
	result,err := menuCollection.Find(context.TODO(),bson.M{})
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured while listing the menu items"})
	}
	defer cancel()
	var allMenus []bson.M

	if err =result.All(ctx ,&allMenus);err !=nil{
		log.Fatal(err)
	}
	c.JSON(http.StatusOK,allMenus )
	}
}
func GetMenu()gin.HandlerFunc{
	return func (ctx *gin.Context)  {
		
	}
}
func CreateMenu()gin.HandlerFunc{
	return func (c *gin.Context)  {
		var ctx , cancel = context.WithTimeout(context.Background(),100*time.Second)
		var menu models.Menu
		 

		if err := c.BindJSON(&menu);err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		validationError:=validate.Struct(menu)
		if validationError!=nil {
			c.JSON(http.StatusBadRequest , gin.H{"error" : validationError.Error()})
			return
		}
		defer cancel()
 		menu.Created_at , _= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		 menu.Updated_at , _= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		 menu.ID = primitive.NewObjectID()
		 menu.Menu_id = menu.ID.Hex()
		
		 result ,insertErr := foodCollection.InsertOne(ctx,menu)
		 if insertErr!=nil {
			msg := fmt.Sprintf("Menu item was not created ")
			c.JSON(http.StatusInternalServerError,gin.H{"err":msg})
			return
		 }
		defer cancel()
		c.JSON(http.StatusOK,result)
		
	}
}
func inTimeSpan(start,end,check time.Time)bool{
	return start.After(time.Now()) &&end.After(start)
}
func UpdateMenu()gin.HandlerFunc{
	return func (c *gin.Context)  {
		var ctx , cancel = context.WithTimeout(context.Background(),100*time.Second)
		var menu models.Menu

		if err := c.BindJSON(&menu);err !=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

		menuId :=c.Param("menu_id")
	 

		var updateObj primitive.D

		if menu.Start_Date !=nil && menu.End_Date!=nil{
			if !inTimeSpan(*menu.Start_Date ,*menu.End_Date,time.Now()) {
				msg:= "kindly retype the time"
				c.JSON(http.StatusInternalServerError,gin.H{"err":msg})
				defer cancel()
				return
			}
		}
		updateObj = append(updateObj, bson.E{"start_date",menu.Start_Date})
		updateObj = append(updateObj, bson.E{"end_date" , menu.End_Date})

		if menu.Name!= "" {
			updateObj =append(updateObj, bson.E{"name" ,menu.Name})
		}
		if menu.Category!= "" {
			updateObj =append(updateObj, bson.E{"category" ,menu.Category})
		}
		menu.Updated_at , _= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"update_at" , menu.Updated_at})

		upsert :=true
		filter := bson.M{"menu_id" :menuId} 
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result , err := menuCollection.UpdateOne(ctx,filter,bson.D{
			"$set",updateObj
			},
			&opt)
			if err!= nil {
				msg := "an error occured update object"
				c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
				return
			}
			defer cancel()
				
	c.JSON(http.StatusOK,result)

	}
	 
}
