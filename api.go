package main

import (

)

// Internal API, create a short api and return it
func CreateShortURL(db DataBase, long string) string {
	if short, ok := db.IsInDb(long); ok {
		return short
	}
	if short, ok := db.Add(long, db.GetId()); ok {
		return "http://54.196.113.135:8080/"+short
	} 
	return ""
}
