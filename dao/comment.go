package dao

import (
	"article-manager/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 增加评论
func (d *Dao) CreateAComment(ctx context.Context, comment *model.Comment) error {
	comment.ID = primitive.NewObjectID()
	_, err := d.ComCol.InsertOne(ctx, comment)
	return err
}

// 根据文章ID获取评论
func (d *Dao) GetCommentsByPaperID(paperID primitive.ObjectID) (comms []*model.Comment, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"paper_id": paperID}

	cur, err := d.ComCol.Find(ctx, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var c *model.Comment
		err = cur.Decode(&c)
		if err != nil {
			return
		}
		comms = append(comms, c)
	}
	return
}

// 根据评论ID获取评论
func (d *Dao) GetCommentByID(commentID primitive.ObjectID) (comm *model.Comment, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	comm = new(model.Comment)
	filter := bson.M{"_id": commentID}
	// 查询并赋予
	err = d.UserCol.FindOne(ctx, filter).Decode(comm)
	if err != nil {
		return nil, err
	}
	return
}

// 根据评论ID删除评论
func (d *Dao) DeleteComment(ctx context.Context, commentID primitive.ObjectID) error {
	filter := bson.M{"_id": commentID}
	_, err := d.ComCol.DeleteOne(ctx, filter)
	return err
}
