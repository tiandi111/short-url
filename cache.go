package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func NewCacheClient() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func CloseCilent(c redis.Conn) error{
	if err := c.Close(); err != nil {
		return err
	}
	return nil
}

func InsertHash(c redis.Conn, pk string, arg *URL) error {
	r, err := c.Do("HMSET", redis.Args{}.Add(pk).AddFlat(arg)...)
	log.Printf("InsertHash: reply[%v]", r)
	if err != nil {
		return err
	}
	return nil
}

func GetHashField(c redis.Conn, pk, field string) (interface{}, error) {
	r, err := c.Do("HGET", pk, field)
	log.Printf("GetHashField: reply[%v]", r)
	if err != nil {
		return nil, err
	}
	return r, err
}

func InsertString(c redis.Conn, key, val string) error {
	r, err := c.Do("SET", key, val)
	if err != nil {
		log.Printf("InsertString: err[%s]\n", err)
		return err
	}
	if r.(string) != "OK" {
		log.Printf("InsertString: SET failed, reply[%s]", r)
	}
	return nil
}

func GetString(c redis.Conn, key string) (string, error) {
	r, err := c.Do("GET", key)
	if err != nil {
		log.Printf("GetString: err[%s]\n", err)
		return "", err
	}
	if r.(string) != "OK" {
		log.Printf("GetString: SET failed, reply[%s]", r)
	}
	return r.(string), nil
}

