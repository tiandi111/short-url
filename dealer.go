package main

var curID int = 0

var idCh = make(chan int, 100)

func GenerateID() int {
	for {
		curID++
		idCh <- curID
	}
}
