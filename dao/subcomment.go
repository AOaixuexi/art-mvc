package dao

import (
	"article-manager/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateASubComment adds a new sub-comment
func (d *Dao) CreateASubComment(ctx context.Context, subComment *model.SubComment) (err error) {
	subComment.ID = primitive.NewObjectID()
	_, err = d.SubcomCol.InsertOne(ctx, subComment)
	return err
}

// GetSubCommentsByCommentID retrieves sub-comments by comment ID
func (d *Dao) GetSubCommentsByCommentID(commentID primitive.ObjectID) (subcoms []*model.SubComment, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"comment_id": commentID}

	cur, err := d.SubcomCol.Find(ctx, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var s *model.SubComment
		err = cur.Decode(&s)
		if err != nil {
			return
		}
		subcoms = append(subcoms, s)
	}
	return
}

// GetSubCommentByID retrieves a sub-comment by its ID
func (d *Dao) GetSubCommentByID(subCommentID primitive.ObjectID) (subcom *model.SubComment, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	subcom = new(model.SubComment)
	filter := bson.M{"_id": subCommentID}
	// 查询并赋予
	err = d.SubcomCol.FindOne(ctx, filter).Decode(subcom)
	if err != nil {
		return nil, err
	}
	return
}

// DeleteSubComment deletes a sub-comment by its ID
func (d *Dao) DeleteSubComment(ctx context.Context, subCommentID primitive.ObjectID) error {
	filter := bson.M{"_id": subCommentID}
	_, err := d.SubcomCol.DeleteOne(ctx, filter)
	return err
}
