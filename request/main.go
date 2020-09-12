package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalf("Cannot connect to NATS server!")
	}

	log.Println("Connected to NATS server.")

	defer nc.Close()

	number := int32(1)

	for {
		payload := fmt.Sprintf("Hello, need reply %d", number)

		response, err := nc.Request("getReply", []byte(payload), 5 * time.Second)

		if err != nil {
			log.Fatalf("Error request: %v", err)
		}

		log.Printf("Response reply: %v", string(response.Data))

		number++
		time.Sleep(1000 * time.Millisecond)
	}
}
