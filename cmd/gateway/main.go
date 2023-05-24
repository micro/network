package main

import (
	"log"
	"os"

	"github.com/micro/network"
)

func main() {
	token := os.Getenv("MICRO_NETWORK_TOKEN")
	if len(token) == 0 {
		log.Fatal("Missing MICRO_NETWORK_TOKEN")
	}

	net := network.New(
		network.WithName("gateway"),
		network.WithToken(token),
	)

	if err := net.Connect(); err != nil {
		log.Fatal(err)
	}
	defer net.Close()
}
