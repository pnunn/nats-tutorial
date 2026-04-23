package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nkeys"
)

func main() {
	// The private seed provided by the user for clientuser
	seed := "SUAIC5POTHANFC66MENDCYMH5UIGG7HQSATVRV2UPILTY4T2G4I3BOIARY"

	// Derive the key pair from the seed
	kp, err := nkeys.FromSeed([]byte(seed))
	if err != nil {
		log.Fatalf("Error deriving key pair from seed: %v", err)
	}

	// Get the public key
	publicKey, err := kp.PublicKey()
	if err != nil {
		log.Fatalf("Error getting public key: %v", err)
	}

	fmt.Printf("Public Key derived from seed '%s': %s\n", seed, publicKey)
	fmt.Println("Expected Public Key in server config: UCEDWC2YVP6N6HYULOBFBOPWYSKW6OKRVQTSP5P3HKSGMEQJHXQBXP2R")
}
