package dao

import (
	"article-manager/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 增加文章
func (d *Dao) CreateAPaper(ctx context.Context, paper *model.Paper) (err error) {
	paper.ID = primitive.NewObjectID()
	_, err = d.PaperCol.InsertOne(ctx, paper)
	return
}

// 通过用户id获取文章
func (d *Dao) GetPapersByUserID(userID primitive.ObjectID) (papers []*model.Paper, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"user_id": userID}

	cur, err := d.PaperCol.Find(ctx, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var p *model.Paper
		err = cur.Decode(&p)
		if err != nil {
			return
		}
		papers = append(papers, p)
	}
	return
}

// 通过文章id和用户id获取文章
func (d *Dao) GetAPaperByIDAndUserID(paperID, userID primitive.ObjectID) (paper *model.Paper, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	paper = new(model.Paper)
	filter := bson.M{"_id": paperID, "user_id": userID}
	// 查询
	err = d.PaperCol.FindOne(ctx, filter).Decode(paper)
	if err != nil {
		return nil, err
	}
	return
}

// 通过文章名和用户id获取文章
func (d *Dao) GetPapersByTitleAndUserID(title string, userID primitive.ObjectID) (papers []*model.Paper, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"title": title, "user_id": userID}

	cur, err := d.PaperCol.Find(ctx, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var p *model.Paper
		err = cur.Decode(&p)
		if err != nil {
			return
		}
		papers = append(papers, p)
	}
	return
}

// 通过文章id更新文章
func (d *Dao) UpdateAPaper(ctx context.Context, paper *model.Paper) (err error) {
	filter := bson.M{"_id": paper.ID}
	update := bson.M{"$set": paper}
	_, err = d.PaperCol.UpdateOne(ctx, filter, update)
	return
}

// 通过文章id删除文章
func (d *Dao) DeleteAPaper(ctx context.Context, paperID primitive.ObjectID) error {
	filter := bson.M{"_id": paperID}
	_, err := d.PaperCol.DeleteOne(ctx, filter)
	return err
}

// 通过用户名字获取用户
func (d *Dao) GetUserByName(name string) (user *model.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user = new(model.User)
	filter := bson.M{"name": name}
	// 查询并赋予
	err = d.UserCol.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return
}

// 通过用户id获取用户
func (d *Dao) GetUserById(id primitive.ObjectID) (user *model.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user = new(model.User)
	filter := bson.M{"_id": id}
	// 查询并赋予
	err = d.UserCol.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return
}

// 通过文章名获取文章
func (d *Dao) GetPapersByTitle(title string) (papers []*model.Paper, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"title": title}

	cur, err := d.PaperCol.Find(ctx, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var p *model.Paper
		err = cur.Decode(&p)
		if err != nil {
			return
		}
		papers = append(papers, p)
	}
	return
}
