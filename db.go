package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type URL struct {
	ID	int
	ID64	string
	LongURL	string
	CreDate	string
	ExpDate	string
	UserID 	int
	TotalClicks	int
	Location	interface{}
}

func NewClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

// 插入一个文件
func InsertOne(ctx context.Context, c *mongo.Collection, document interface{}) error {
	_, err := c.InsertOne(context.Background(), document)
	if err != nil {
		log.Println("Insert failed:", err)
		return err
	}
	return nil
}

//查询一个文件
func FindOne(ctx context.Context, c *mongo.Collection, filter interface{}) (bson.Raw, bool) {
	ret, err := c.FindOne(ctx, filter).DecodeBytes()
	if err != nil {
		log.Println("Find failed:", err)
		return nil, false
	}
	return ret, true
}

