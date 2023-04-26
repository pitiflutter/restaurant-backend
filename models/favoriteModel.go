package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Favorite struct{
	ID primitive.ObjectID `bson:"_id"`
	User_id *string `json:"user_id" validate:"required"`
	Food_id *string `json:"food_id" validate:"required"`
}