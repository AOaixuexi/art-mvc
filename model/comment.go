package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// comment model
type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserName  string             `bson:"username" json:"username"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at"`
	PaperID   primitive.ObjectID `bson:"paper_id" json:"-"`
	UserID    primitive.ObjectID `bson:"user_id" json:"-"`
}

func (Comment) CommentTableName() string {
	return "comment"
}
