package main

import "math/rand"

var idCh = make(chan int32, 100)

func GenerateID() int {
	for {
		idCh <- rand.Int31()
	}
}
