package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant/database"
 	"restaurant/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var favoriteCollection *mongo.Collection = database.OpenCollection(database.Client, "favorite")

func GetFavorites() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Param("user_id")
		matchStage := bson.D{{"$match", bson.D{{"user_id", user_id}}}}
		cursor, err := favoriteCollection.Aggregate(context.TODO(), mongo.Pipeline{matchStage})
		if err != nil {
			panic(err)
		}
		var results []models.Favorite
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}
		for _, result := range results {
			fmt.Println(*&result.Favorite_id)
		}
		c.JSON(http.StatusOK, results)
	}
}
func DeleteFavorite() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordPrePage, err := strconv.Atoi(c.Query("recordPerPage"))

		if err != nil || recordPrePage < 1 {
			recordPrePage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPrePage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"data", bson.D{{"$push", "$$ROOT"}}}}}}

		projectStage := bson.D{
			{
				"$project", bson.D{
					{"_id", 0},
					{"total_count", 1},
					{"food_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPrePage}}}},
				}}}

		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "error occured while listing food items"})

		}
		var allFoods []bson.M
		if err = result.All(ctx, &allFoods); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"results": allFoods[0]})
	}
}
func CreateFavorite() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var favorite models.Favorite
 		 

		if err := c.BindJSON(&favorite); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(favorite)
		validationError := validate.Struct(favorite)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
			return
		}
 		favorite.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		favorite.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		favorite.ID = primitive.NewObjectID()
		favorite.Favorite_id = favorite.ID.Hex()
	 

		result, insertErr := favoriteCollection.InsertOne(ctx, favorite)
		if insertErr != nil {
			msg := fmt.Sprintf("Favorite item was not created ")
			c.JSON(http.StatusInternalServerError, gin.H{"err": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}
