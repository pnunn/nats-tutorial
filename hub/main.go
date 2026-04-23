package main

import (
	"encoding/json"
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

	// 1. Subscribe to Feedback from Leaf
	_, err = js.Subscribe("feedback.leaf.hub", func(msg *nats.Msg) {
		var feedback common.FeedbackMsg
		if err := json.Unmarshal(msg.Data, &feedback); err != nil {
			log.Printf("Feedback Unmarshal Error: %v", err)
			return
		}
		log.Printf("HUB RECEIVED FEEDBACK: %+v", feedback)
		msg.Ack()
	}, nats.Durable("hub-feedback-consumer"), nats.ManualAck())
	if err != nil {
		log.Fatal(err)
	}

	// 2. Subscribe to Watchdog from Leaf
	_, err = js.Subscribe("watchdog.leaf", func(msg *nats.Msg) {
		var watchdog common.WatchdogMsg
		if err := json.Unmarshal(msg.Data, &watchdog); err != nil {
			log.Printf("Watchdog Unmarshal Error: %v", err)
			return
		}
		log.Printf("HUB RECEIVED WATCHDOG: %+v", watchdog)
		msg.Ack()
	}, nats.Durable("hub-watchdog-consumer"), nats.ManualAck())
	if err != nil {
		log.Fatal(err)
	}

	// 3. Periodically Publish Dispatch to Leaf
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		dispatch := common.DispatchMsg{
			Timestamp:      time.Now(),
			SettlementDate: time.Now().Format("2006-01-02"),
			DispatchFlag:   true,
			DispatchCap:    150.5,
		}
		data, _ := json.Marshal(dispatch)
		_, err := js.Publish("dispatch.hub.leaf", data)
		if err != nil {
			log.Printf("Publish Dispatch Error: %v", err)
		} else {
			log.Printf("HUB PUBLISHED DISPATCH: %+v", dispatch)
		}
	}
}
