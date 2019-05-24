package main

import (
	"context"
	"github.com/shorturl/short-url/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Create a short url document
func CreateShortURL(ctx context.Context, c *mongo.Collection, long string) (string, error) {
	// if there's a short url for the given long url
	// return it
	if short := FindShortURL(ctx, c, bson.D{{"longurl", long}}); len(short) != 0 {
		return "http://127.0.0.1:8080/"+short,nil
	}
	// otherwise, create a new record
	id := <-IdCh
	short := util.Encode(id)
	document := URL{
		id,
		short,
		long,
		time.Now().Format("201905221602"),
		1,
		"",
		0,
		0,
		0,
	}
	err := InsertOne(ctx, c, document)
	if err != nil {
		return "", err
	}
	return DomainName+short, err
}

// Find long url according to given filter
func FindLongURL(ctx context.Context, c *mongo.Collection, filter interface{}) string {
	ret, ok := FindOne(ctx, c, filter)
	if !ok {
		return ""
	}
	return string(ret.Lookup("longurl").StringValue())
}

// Find short url according to given filter
func FindShortURL(ctx context.Context, c *mongo.Collection, filter interface{}) string {
	ret, ok := FindOne(ctx, c, filter)
	if !ok {
		return ""
	}
	return string(ret.Lookup("id64").StringValue())
}