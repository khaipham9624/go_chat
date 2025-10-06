package main

import (
	"gochat/src/server/restserver"
)

func main() {
	// db.Init()
	// db.CreateUserIndex()
	restserver.Start(8080)
}
