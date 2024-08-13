package dao

import (
	"article-manager/conf"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dao struct {
	Mongoclient *mongo.Client
	UserCol     *mongo.Collection
	PaperCol    *mongo.Collection
	ComCol      *mongo.Collection
	SubcomCol   *mongo.Collection
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		Mongoclient: NewMongo(c.Mongo),
	}

	d.UserCol = d.Mongoclient.Database("forumDevDB").Collection("user")
	d.PaperCol = d.Mongoclient.Database("forumDevDB").Collection("paper")
	d.ComCol = d.Mongoclient.Database("forumDevDB").Collection("comment")
	d.SubcomCol = d.Mongoclient.Database("forumDevDB").Collection("subcomment")
	return

}

func NewMongo(c *conf.Mongo) (client *mongo.Client) {
	var err error
	clientOptions := options.Client()
	if c.Username != "" && c.Password != "" {
		clientOptions.SetAuth(options.Credential{Username: c.Username, Password: c.Password})
	}

	clientOptions.SetHosts(c.Addrs)
	clientOptions.SetMaxPoolSize(c.MaxPool)

	// 连接
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("mongo.Connect error %v", err)
		return
	}

	// 测试连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("client.Ping error %v", err)
	}
	return
}
