package main

import (
	//"fmt"
	//"strconv"
)

// DataBase interface encapsulate the 
type DataBase interface {
	IsInDb (url string) (string, bool)
	GetLongURL (url string) string
	GetId () int
	Add (long string, id int) (string, bool)
}

// DbDriver provide an access to database
// Now, we use two maps and a integer variable ID to mock a database
// name    key     value
// StoL	   id	   long url
// LtoS    longurl  id
type DbDriver struct {
	ID	int
	StoL	map[string]string
	LtoS	map[string]string
}

// Initialize a new database
func NewDataBase() DataBase {
	id := -1
	stol := make(map[string]string, 1)
	ltos := make(map[string]string, 1)
	// Here we should return the pointer instead of struct itself
	// Becasue the method required by DataBase interface is attached
	// to the pointer of DbDrive
	// In other words, it is the pointer of DbDriver a Database 
	// not DbDriver itself a DataBase!
	return &DbDriver{ id, stol, ltos}
}

// IsInDb check if the given long url is already in database
// If so, return (shorturl, true)
// otherwise, return ("", false)
func (db *DbDriver) IsInDb(url string) (string, bool){
	if short, ok := db.LtoS[url]; ok {
		return short, true
	}
	return "", false
}

// GetLongURL return the corresponding long url
func (db *DbDriver) GetLongURL(url string) string {
	if long, ok := db.StoL[url]; ok {
		return long
	}
	return ""
}

// GetId get an ID from database
// Then the ID is uesd to create a new long-short url pair in database
func (db *DbDriver) GetId() int {
	db.ID++
	return db.ID
}

// Add the given long url to the database
func (db *DbDriver) Add(long string, id int) (string, bool) {
	short := Encode(id)
	db.StoL[short] = long
	db.LtoS[long] = short
	return short, true
}

//func Encode(id int) string {
//	return strconv.Itoa(id)
//}



