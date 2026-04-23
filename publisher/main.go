package main

import (
	"log"
	"nats-tutorial/common"
	"time"
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

	subject := "foo.test"
	for i := 1; i <= 5; i++ {
		msg := "Message " + string(rune('0'+i))
		ack, err := js.Publish(subject, []byte(msg))
		if err != nil {
			log.Printf("Publish Error: %v", err)
			continue
		}
		log.Printf("Published to [%s], Stream: %s, Sequence: %d", subject, ack.Stream, ack.Sequence)
		time.Sleep(1 * time.Second)
	}
}
