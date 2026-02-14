package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email        string             `bson:"email" json:"email"`
	PasswordHash string             `bson::"Passwordhash"`
	Role         string             `bson:"role" json:"role"`
	CreatedAt    time.Time          `bson:"createdAt" json"createdAt"`
	UpdatedAt    time.Time          `bson:"updateddAt" json"updatedAt"`
}
type PublicUser struct{	
	ID string `json:"id"`
	Email string `json:"email"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"id"`
	UpdatedAt time.Time `json:"id"`
}
func ToPublic(u User) PublicUser{
	return PublicUser{
		ID: u.ID.Hex(),
		Email: u.Email,
		Role: u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
