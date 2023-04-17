package main

import (
	"encoding/json"
	"github.com/RokibulHasan7/notes/nats/practice/jetstream/model"
	"github.com/nats-io/nats.go"
	"log"
)

func consumeReviews(js nats.JetStreamContext) {
	_, err := js.Subscribe(SubjectNameReviewCreated, func(m *nats.Msg) {
		err := m.Ack()
		log.Println("Message Ack: ", m.Ack())
		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}
		var review model.Review
		err = json.Unmarshal(m.Data, &review)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Consumer  =>  Subject: %s  -  ID:%s  -  Author: %s  -  Rating:%d\n", m.Subject, review.Id, review.Author, review.Rating)

		// send answer via JetStream using another subject if you need
		// js.Publish(config.SubjectNameReviewAnswered, []byte(review.Id))
	})

	if err != nil {
		log.Println(err.Error())
		log.Println("Subscribe failed")
		return
	}
}
