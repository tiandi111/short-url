package main

import "math/rand"

var IdCh = make(chan int32, 100)

func GenerateID() int {
	for {
		IdCh <- rand.Int31()
	}
}
