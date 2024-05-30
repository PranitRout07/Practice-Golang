package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MovieName string             `json:"moviename,omitempty"`
	Duration  int                `json:"duration,omitempty"`
}
