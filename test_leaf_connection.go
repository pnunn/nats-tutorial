package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	url := "nats://127.0.0.1:4223"
	log.Println("Attempting to connect to:", url)

	opts := []nats.Option{
		nats.Timeout(10 * time.Second), // Generous timeout
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("Disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("Reconnected to %s", nc.ConnectedUrl())
		}),
	}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		log.Fatalf("FATAL: Connection failed: %v", err)
	}
	defer nc.Close()

	log.Println("SUCCESS: Successfully connected to NATS server!")
	
	// Correctly handle the two return values from RTT()
	rtt, err := nc.RTT()
	if err != nil {
		log.Printf("Could not get RTT: %v", err)
	} else {
		log.Println("Server RTT:", rtt)
	}
}
