package main

import (
	"encoding/json"
	"log"
	"nats-tutorial/common"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	url := common.GetConfig("NATS_LEAF_SERVER", "nats://127.0.0.1:4223")
	user := common.GetConfig("NATS_USER", "clientuser")
	password := common.GetConfig("NATS_PASSWORD", "clIent2026Pw")

	opts := []nats.Option{
		nats.UserInfo(user, password),
		nats.Timeout(10 * time.Second),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(5 * time.Second),
	}
	
	nc, err := nats.Connect(url, opts...)
	if err != nil {
		log.Fatalf("Connect Error: %v", err)
	}
	defer nc.Close()

	log.Printf("Connected to NATS at: %s", nc.ConnectedAddr())

	// Initialize JetStream with the 'hub' domain to reach the Hub cluster
	js, err := nc.JetStream(nats.Domain("hub"))
	if err != nil {
		log.Fatalf("JetStream Error: %v", err)
	}
	log.Println("JetStream context initialized via 'hub' domain")

	// 1. Subscribe to Dispatch from Hub
	_, err = js.Subscribe("dispatch.hub.leaf", func(msg *nats.Msg) {
		var dispatch common.DispatchMsg
		if err := json.Unmarshal(msg.Data, &dispatch); err != nil {
			log.Printf("Dispatch Unmarshal Error: %v", err)
			return
		}
		log.Printf("LEAF RECEIVED DISPATCH: %+v", dispatch)

		// 2. Publish Feedback Immediately
		feedback := common.FeedbackMsg{
			Timestamp:            time.Now(),
			ActivePower:          148.2,
			DispatchCapFeedback:  dispatch.DispatchCap,
			DispatchFlagFeedback: dispatch.DispatchFlag,
		}
		data, _ := json.Marshal(feedback)
		
		subject := "feedback.leaf.hub"
		_, err := js.Publish(subject, data)
		if err != nil {
			log.Printf("Publish Feedback Error: %v", err)
		} else {
			log.Printf("LEAF PUBLISHED FEEDBACK to %s", subject)
		}

		msg.Ack()
	}, nats.Durable("leaf-dispatch-consumer"), nats.ManualAck())

	if err != nil {
		log.Fatalf("Failed to subscribe to dispatch stream: %v", err)
	}
	log.Println("Successfully subscribed to dispatch.hub.leaf")

	// 3. Periodically Publish Watchdog
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			watchdog := common.WatchdogMsg{
				Timestamp: time.Now(),
				Source:    "Leaf Node 1",
			}
			data, _ := json.Marshal(watchdog)

			subject := "watchdog.leaf"
			_, err := js.Publish(subject, data)
			if err != nil {
				log.Printf("Publish Watchdog Error: %v", err)
			} else {
				log.Printf("LEAF PUBLISHED WATCHDOG to %s", subject)
			}
		}
	}()

	// Keep alive
	select {}
}
