package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"firstname" validate:"required, min=2 max=100"`
	Last_name     *string            `json:"lastname" validate:"required, min=2 max=100"`
	Password      *string            `json:"Password" validate:"required, min=6 max=100"`
	Email         *string            `json:"email" validate:"email,required"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
