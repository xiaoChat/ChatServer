package main

import (
	"xiaochat/app"
	"xiaochat/app/db"
)

func main() {
	// init db
	db.Init()
	// run api server
	app.StartServer()
}
