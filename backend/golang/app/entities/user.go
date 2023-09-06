package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Users struct {
		ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
		Name      string             `bson:"name" json:"name"`
		Email     string             `bson:"email" json:"email"`
		Password  string             `bson:"password" json:"password"`
		CreatedAt time.Time          `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	}
	UsersCreateInput struct {
		Name     string `bson:"name" json:"name"`
		Email    string `bson:"email" json:"email"`
		Password string `bson:"password" json:"password"`
	}
)

func (Users) UsersCollection() string {
	return "users"
}
