package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

func main() {
	seed := "SUAIC5POTHANFC66MENDCYMH5UIGG7HQSATVRV2UPILTY4T2G4I3BOIARY"
	url := "tls://natsgw.marketdispatch.com.au:4222"

	opt, _ := nats.NkeyOptionFromSeed(seed)

	// NO retries to see the raw error immediately
	opts := []nats.Option{
		opt,
		nats.Timeout(2 * time.Second),
	}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		log.Fatalf("Raw Error connecting: %v", err)
	}
	defer nc.Close()

	fmt.Println("Successfully connected!")
}