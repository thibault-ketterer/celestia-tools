package main

import (
	"encoding/hex"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"fmt"
	"math/rand"
    	"os"
	"strconv"
	"net/http"
)

// ATTENTION:
// In order to generate a random namespace ID and hex-encoded data, please enter a valid integer on Line 17 after the `:=`.
//
// Example: `seed := 5405`


func makeRequest(url string, namespaceID string, gasLimit int, fee int, data string) ([]byte, error) {
    requestData := map[string]interface{}{
        "namespace_id": namespaceID,
        "gas_limit":    gasLimit,
        "fee":          fee,
        "data":         data,
    }
    jsonBody, err := json.Marshal(requestData)
    if err != nil {
        return nil, fmt.Errorf("Error marshaling JSON data: %s", err)
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
    if err != nil {
        return nil, fmt.Errorf("Error creating POST request: %s", err)
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("Error making POST request: %s", err)
    }
    defer resp.Body.Close()

    // Read response body
    respBodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    // Print response body to console
    fmt.Println(string(respBodyBytes))

    // Return response body
    return respBodyBytes, nil
}

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
		fmt.Printf("calling request to\n '{\"namespace_id\": \"%s\",\n \"data\": \"%s\",\n \"gas_limit\": 80000,\n \"fee\": 2000}'\n http://%s:26659/submit_pfb\n", nID, msg, IP)

		// responseBody, err := makeRequest("http://celestia.lankou.org:26659/submit_pfb", nID, 80000, 2000, msg)
		responseBody, err := makeRequest(fmt.Sprintf("http://%s:26659/submit_pfb", IP), nID, 80000, 2000, msg)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}

		fmt.Println(string(responseBody))
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
