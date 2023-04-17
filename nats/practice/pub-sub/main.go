package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	// Connect to the NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe to a topic
	sub, err := nc.SubscribeSync("updates")
	if err != nil {
		log.Fatal(err)
	}

	// Start a goroutine to handle incoming messages
	go func() {
		for {
			msg, err := sub.NextMsg(time.Second * 10)
			if err != nil {
				if nats.ErrTimeout == err {
					continue
				}
				log.Fatal(err)
			}

			fmt.Printf("Received a message: %s\n", string(msg.Data))
		}
	}()

	// Publish some messages
	for i := 1; i <= 10; i++ {
		err := nc.Publish("updates", []byte(fmt.Sprintf("Message %d", i)))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Published message %d\n", i)
	}

	// Wait for messages to be received
	time.Sleep(time.Second * 1)
}
