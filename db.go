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

// 查找一个长url是否存在。若存在，返回该文件；否则返回nil
func FindLongURL(c *mongo.Collection, ctx context.Context, id64 string) string {
	filter := bson.M{"ID64": id64}
	ret, ok := FindOne(c, ctx, filter)
	if !ok {
		return ""
	}
	//转化为string
	return string(ret.Lookup("LongURL").Value[:])
}

func FindShortURL(c *mongo.Collection, ctx context.Context, long string) string {
	filter := bson.M{"LongURL": long}
	ret, ok := FindOne(c, ctx, filter)
	if !ok {
		return ""
	}
	//转化为string
	return string(ret.Lookup("ID64").Value[:])
}

// 插入一个文件
func InsertOne(c *mongo.Collection, ctx context.Context, document interface{}) error {
	_, err := c.InsertOne(context.Background(), document)
	if err != nil {
		log.Println("Insert failed:", err)
		return err
	}
	return nil
}

//查询一个文件
func FindOne(c *mongo.Collection, ctx context.Context, filter interface{}) (bson.Raw, bool) {
	ret, err := c.FindOne(ctx, filter).DecodeBytes()
	if err != nil {
		log.Println("Find failed:", err)
		return nil, false
	}
	return ret, true
}

