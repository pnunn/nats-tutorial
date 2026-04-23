package main

import (
	"log"
	"nats-tutorial/common"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := common.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	streamName := "test-stream"
	consumerName := "my-pull-consumer"
	subject := "foo.test"

	sub, err := js.PullSubscribe(subject, consumerName, nats.BindStream(streamName))
	if err != nil {
		log.Fatal(err)
	}

	for {
		msgs, err := sub.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				log.Println("Timeout waiting for messages...")
				continue
			}
			log.Printf("Fetch Error: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, msg := range msgs {
			log.Printf("Received: %s", string(msg.Data))
			if err := msg.Ack(); err != nil {
				log.Printf("Ack Error: %v", err)
			} else {
				log.Println("Acknowledged message")
			}
		}
	}
}
