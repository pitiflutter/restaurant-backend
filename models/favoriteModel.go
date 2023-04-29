package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Favorite struct {
	ID          primitive.ObjectID `bson:"_id"`
	Favorite_id string             `json:"favorite_id"`
	User_id     *string            `json:"user_id" validate:"required"`
	Food_id     *string            `json:"food_id" validate:"required"`
	Food_detail Food               `json:"food_detail"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"update_at"`
}
