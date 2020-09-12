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

	nc.Subscribe("getReply", func(m *nats.Msg) {
		log.Printf("Received a request: %s\n", string(m.Data))

		reply := "You're OK!"
		m.Respond([]byte(reply))

		log.Printf("Reply sent.")
	})

	nc.Flush()

	runtime.Goexit()
}
