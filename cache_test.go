package main

import (
	//"fmt"
	"testing"
)


func TestNew(t *testing.T) {
	if _, ok := NewDataBase().(DataBase); !ok {
		t.Errorf("NewDataBase should return a DataBase type!")
	}
}

func TestGetId(t *testing.T) {
	db := NewDataBase()
	for i := 0; i < 100; i++ {
		id := db.GetId()
		if id != i {
			t.Errorf("The %dth operation is incorrect, expect %d, but %d", i, i, id)
		}
	}
}

func TestDataBase(t *testing.T) {
	db := NewDataBase()
	urls := []string {"abc", "edf", "ghi"}
	exp := []string {"0", "1", "2"}
	for _, url := range urls {
		db.Add( url, db.GetId() )
	}
	for i, url := range urls {
		if res, _ := db.IsInDb( url ); res != exp[i] {
			t.Errorf("The %dth operation is incorrect, expect %v, but %v", i, exp[i], res)
		}
	}
	if _, ok := db.IsInDb( "addfds" ); ok {
		t.Errorf("The 3th operation is incorrect, expect false, but true")	
	}
}
