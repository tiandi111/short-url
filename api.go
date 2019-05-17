package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// Internal API, create a short url and return it
func CreateShortURL(c *mongo.Collection, ctx context.Context, long string) (string, error) {
	//if short, ok := c.IsInDb(long); ok {
	//	return short
	//}
	//if short, ok := db.Add(long, db.GetId()); ok {
	//	return "http://54.196.113.135:8080/"+short
	//}
	id := <- idCh
	short := Encode(id)
	document := URL{
		id,
		short,
		long,
		"",
		"",
		0,
		0,
		0,
	}
	err := InsertOne(c, ctx, document)
	return "http://54.196.113.135:8080/"+short, err
}
