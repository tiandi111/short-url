package main

import (
	"fmt"
)

type DbDriver struct {
	ID	int
	short	map[int][string]
	long	map[string][string]
}

func (db *DbDriver) NewDatabase() {
	id := 0
	short := make(map[int][string], 1)
	long  := make(map[string][string], 1)
}

func (db *DbDriver) IsInDb(url string) (string bool){
	
}
