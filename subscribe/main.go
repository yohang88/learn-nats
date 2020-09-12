package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalf("Cannot connect to NATS server!")
	}

	log.Println("Connected to NATS server.")

	// Simple Async Subscriber
	nc.Subscribe("hello", func(m *nats.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
	})

	nc.Flush()

	runtime.Goexit()
}
