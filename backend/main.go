package main

import (
	"log"

	"github.com/robert-inkpen/randeng_co_mirror/backend/config"
	v1 "github.com/robert-inkpen/randeng_co_mirror/backend/v1"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		panic(err)
	}

	server := v1.NewServer(cfg)

	if err := server.Start(); err != nil {
		log.Fatalf("Failed binding server: %v", err)
	}
}
