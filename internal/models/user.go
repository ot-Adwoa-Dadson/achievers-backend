package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FullName   string             `bson:"fullName"`
	Email      string             `bson:"email"`
	Password   string             `bson:"password"`
	Role       string             `bson:"role"` // ADMIN or LEADER
	SeniorCell string             `bson:"seniorCell,omitempty"`
	CreatedAt  time.Time          `bson:"createdAt"`
}
