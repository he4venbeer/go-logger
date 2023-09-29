package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestLog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Method    string             `json:"method" bson:"method"`
	URI       string             `json:"uri" bson:"uri"`
	CreatedAt time.Time          `json:"createAt" bson:"created_at"`
}
