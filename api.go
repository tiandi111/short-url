package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Internal API, create a short url and return it
func CreateShortURL(ctx context.Context, c *mongo.Collection, long string) (string, error) {
	//如果该长URL已有对应短URL存在，则返回短URL
	//否则创建新的长-短URL对
	if short := FindShortURL(ctx, c, bson.D{{"longurl", long}}); len(short) != 0 {
		return short,nil
	}
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
	err := InsertOne(ctx, c, document)
	return "http://127.0.0.1:8080/"+short, err
}

//查找长URL
func FindLongURL(ctx context.Context, c *mongo.Collection, filter interface{}) string {
	ret, ok := FindOne(ctx, c, filter)
	if !ok {
		return ""
	}
	//转化为string
	//return string(ret.Lookup("longurl").Value[:])
	return rawValueToString(ret.Lookup("longurl").Value)
}

//查找短URL
func FindShortURL(ctx context.Context, c *mongo.Collection, filter interface{}) string {
	ret, ok := FindOne(ctx, c, filter)
	if !ok {
		return ""
	}
	//转化为string
	return rawValueToString(ret.Lookup("id64").Value)
}

// 使用Lookup时，mongodb的go驱动返回一个[]byte，这个
// []byte在原始数据之前添加了3个byte，之后添加一个byte
// 注意：只有string格式的数据会有这种现象，所以该函数
// 只适用于格式为string的数据
func rawValueToString(r []byte) string {
	l := len(r)
	return string(r[4:l-1])
}