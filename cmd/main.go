package main

import (
	"log"

	"gochat-witai/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
