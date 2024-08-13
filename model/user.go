package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// user model
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Password string             `bson:"password" json:"password"`
	Papers   []Paper            `bson:"papers,omitempty" json:"papers"`
	Comments []Comment          `bson:"comments,omitempty" json:"comments"`
}

func (User) UserTableName() string {
	return "user"
}
