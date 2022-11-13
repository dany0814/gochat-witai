package main

import (
	"log"

	"github.com/dany0814/gochat-witai/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
