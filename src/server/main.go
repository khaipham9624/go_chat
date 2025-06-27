package main

import (
	"gochat/src/server/restserver"
)

func main() {
	restserver.Start(8080)
}
