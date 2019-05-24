package main

type URL struct {
	ID	int32			// id
	ID64	string		// 64base of id
	LongURL	string
	CreDate	string
	Duration int
	ExpDate	string
	UserID 	int
	TotalClicks	int
	Location	interface{}
}
