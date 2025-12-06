package main

import (
	"mkassel.com/task-manager/internal/server"
)

func main() {
	srv := server.New(":8080")
	srv.Start()
}
