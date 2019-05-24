package main

import "github.com/gomodule/redigo/redis"

func NewCacheClient() (redis.Conn, error) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return nil, err
	}
	return c, nil
}

func CloseCilent(c redis.Conn) error{
	if err := c.Close(); err != nil {
		return err
	}
	return nil
}

func FindOneInCache() {

}
