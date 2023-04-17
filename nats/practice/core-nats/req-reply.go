package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalln(err)
	}

	defer nc.Close()
	//respond(nc)
	sync(nc)
}

func respond(nc *nats.Conn) {
	_, err := nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("Request received:", string(msg.Data))

		err := msg.Respond([]byte("Here you go!"))
		if err != nil {
			// `message does not have a reply` error will be shown for Publish() call when it's a Publish instead of request
			fmt.Printf("error on Respond() %s \n", err)
			return
		}
	})

	// this returns error because the sub "foo" replies to inbox but here
	// we are publishing instead of requesting
	if err := nc.Publish("foo", []byte("Message")); err != nil {
		log.Fatalln(err)
	}

	reply, err := nc.Request("foo", []byte("Give me data"), 4*time.Second)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Got Reply:", string(reply.Data))
}

func sync(nc *nats.Conn) {
	sub, err := nc.SubscribeSync("bar")
	if err != nil {
		log.Fatal(err)
	}

	nc.Request("bar", []byte("hello universe"), time.Second)
	fmt.Println("request was called.")

	// Send the request, If processing is synchronous, use Request() which returns the response message.
	// PublishRequest is similar to Publish, Difference is,
	// It expects a response on the reply subject. whereas Publish waits for the response.
	nc.PublishRequest("sub", "bar", []byte("hello world")) // `no responders available for request` error will be thrown
	nc.Flush()

	// Wait for a single response
	for {
		msg, err := sub.NextMsg(3 * time.Second)
		if err != nil {
			log.Fatal(err)
		}

		response := string(msg.Data)
		fmt.Println(response)
		break
	}
	sub.Unsubscribe()
}
