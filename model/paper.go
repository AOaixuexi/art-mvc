package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// paper model
type Paper struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title      string             `bson:"title" json:"title"`
	AuthorName string             `bson:"author_name" json:"author_name"`
	Content    string             `bson:"content" json:"content"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	UserID     primitive.ObjectID `bson:"user_id" json:"-"`
	Comments   []Comment          `bson:"comments,omitempty" json:"comments"`
}

func (Paper) PaperTableName() string {
	return "paper"
}
