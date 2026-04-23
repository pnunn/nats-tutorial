package main

import (
	"fmt"
	"os"

	"github.com/nats-io/nkeys"
)

func main() {
	seed := "SUAIC5POTHANFC66MENDCYMH5UIGG7HQSATVRV2UPILTY4T2G4I3BOIARY"
	kp, err := nkeys.FromSeed([]byte(seed))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	pub, _ := kp.PublicKey()
	fmt.Println("Public Key for seed in .env:", pub)
}
