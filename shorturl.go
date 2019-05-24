package main

type URL struct {
	ID	int32			`redis:"id"`// id
	ID64	string		`redis:"id64"`// 64base of id
	LongURL	string		`redis:"longurl"`
	CreDate	string		`redis:"credate"`
	Duration int		`redis:"duration"`
	ExpDate	string		`redis:"expdate"`
	UserID 	int			`redis:"userid"`
	TotalClicks	int		`redis:"totalclicks"`
	Location	interface{}
}
