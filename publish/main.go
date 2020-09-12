package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalf("Cannot connect to NATS server!")
	}

	log.Println("Connected to NATS server.")

	defer nc.Close()

	// Simple Publisher
	nc.Publish("hello", []byte("Hello World"))
	nc.Flush()
}
