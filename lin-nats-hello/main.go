package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to NATS:", err)
		return
	}
	defer nc.Close()
	// Publish a message
	nc.Publish("hello", []byte("Hello from NATS!"))
	// Subscribe to a topic
	nc.Subscribe("hello", func(m *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(m.Data))
	})
	select {}
}
