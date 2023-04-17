package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	sy "sync"
)

func ttt() {

	ns, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer ns.Close()

	subject := "test.sub"

	wg := sy.WaitGroup{}
	wg.Add(1)

	_, err = ns.Subscribe(subject, func(m *nats.Msg) {
		fmt.Printf("Received: %s\n", string(m.Data))
		wg.Done()
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Subscribed to ", subject)

	wg.Wait()

}
