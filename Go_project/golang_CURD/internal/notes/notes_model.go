package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Pinned    bool               `bson:"pin" json:"pin"`
	CreatedAt time.Time          `bson:"CreatedAt" json:"CreatedAt"`
	UpdatedAt time.Time          `bson:"UpdatedAt" json:"UpdatedAt"`
}

type CreateNoteRequest struct {
	Title   string `bson:"title" binding:"required"`
	Content string `bson:"content" binding:"required"`
	Pinned  bool   `bson:"pin" `
}

type UpdateNoteReq struct {
	Title   string `bson:"title" binding:"required"`
	Content string `bson:"content" binding:"required"`
	Pinned  bool   `bson:"pin" `
}
