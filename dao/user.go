package dao

import (
	"article-manager/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *Dao) CreateAUser(ctx context.Context, user *model.User) (err error) {
	user.ID = primitive.NewObjectID()
	_, err = d.UserCol.InsertOne(ctx, user)
	return
}

func (d *Dao) GetAUser(name string, password string) (user *model.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user = new(model.User)
	filter := bson.M{"name": name, "password": password}
	// 查询
	err = d.UserCol.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return
}
