package main

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/shorturl/short-url/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// Create a short url document
func CreateShortURL(ctx context.Context, c *mongo.Collection, cr redis.Conn, long string) (string, error) {
	// if there's a short url for the given long url
	// return it, first check cache then database
	if reply, _ := GetHashField(cr, long, "id64"); reply != nil {
		return DomainName+string(reply.([]byte)), nil
	}
	if short := FindShortURL(ctx, c, cr, bson.D{{"longurl", long}}); len(short) != 0 {
		return DomainName+short,nil
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

	// If insert to database succeed, try insert to cache
	err = InsertHash(cr, long, &document)
	if err != nil {
		return "", err
	}

	return DomainName+short, err
}

// Find long url according to given filter
func FindLongURL(ctx context.Context, c *mongo.Collection, cr redis.Conn, filter interface{}) string {
	ret, ok := FindOne(ctx, c, filter)
	if !ok {
		return ""
	}
	return string(ret.Lookup("longurl").StringValue())
}

// Find short url according to given filter
func FindShortURL(ctx context.Context, c *mongo.Collection, cr redis.Conn, filter interface{}) string {
	ret, ok := FindOne(ctx, c, filter)
	if !ok {
		return ""
	}
	return string(ret.Lookup("id64").StringValue())
}