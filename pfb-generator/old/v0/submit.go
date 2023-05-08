package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
    	"os"
	"strconv"
)

// ATTENTION:
// In order to generate a random namespace ID and hex-encoded data, please enter a valid integer on Line 17 after the `:=`.
//
// Example: `seed := 5405`

func main() {
	//var s int;
	if len(os.Args) < 3 {
		fmt.Printf("\n" +
		"usage\n\n" +
		"	go run submit.go <seed> <NODEIP> | sh\n"+
		"like this\n\n"+
		"	go run submit.go 1234 1.1.1.1 | sh\n\n")
		os.Exit(1)
	}
	if seed, err := strconv.Atoi(os.Args[1]); err == nil {
		// fmt.Printf("seed [%d]\n", seed)

		rand.Seed(int64(seed))

		// generate a random namespace ID
		nID := generateRandHexEncodedNamespaceID()

		// generate a random hex-encoded message
		msg := generateRandMessage()

		IP := os.Args[2]
		//fmt.Println(fmt.Sprintf("My hex-encoded namespace ID: %s\n\nMy hex-encoded message: %s", nID, msg))
		//fmt.Println(fmt.Sprintf("namespace=\"%s\"\ndata=\"%s\"\n", nID, msg))
		fmt.Printf("curl -X POST -d '{\"namespace_id\": \"%s\", \"data\": \"%s\", \"gas_limit\": 80000, \"fee\": 2000}' http://%s:26659/submit_pfb\n", nID, msg, IP)

	}
}

// generateRandHexEncodedNamespaceID generates 8 random bytes and
// returns them as a hex-encoded string.
func generateRandHexEncodedNamespaceID() string {
	nID := make([]byte, 8)
	_, err := rand.Read(nID)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(nID)
}

// generateRandMessage generates a message of an arbitrary length (up to 100 bytes)
// and returns it as a hex-encoded string.
func generateRandMessage() string {
	lenMsg := rand.Intn(100)
	msg := make([]byte, lenMsg)
	_, err := rand.Read(msg)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(msg)
}
