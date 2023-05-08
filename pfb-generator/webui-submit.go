package main


import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

// type Response struct {
// 	Message string `json:"message"`
// }


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


func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // w.Header().Set("Access-Control-Allow-Origin", "*")

        if r.Method == "POST" {
	    body, err := ioutil.ReadAll(r.Body)
            if err != nil {
                http.Error(w, "Bad Request", http.StatusBadRequest)
                return
            }

	    // err := r.ParseForm()
            // if err != nil {
            //     http.Error(w, "Failed to parse form", http.StatusBadRequest)
            //     return
            // }

	    // Parse the payload as JSON
            var payload map[string]interface{}
            err = json.Unmarshal(body, &payload)
            if err != nil {
                http.Error(w, "Bad Request", http.StatusBadRequest)
                return
            }
	    
            url := payload["url"].(string)

            fee, err := strconv.Atoi(payload["fee"].(string))
            if err != nil {
                http.Error(w, "Failed to parse fee", http.StatusBadRequest)
                return
            }

            gasLimit, err := strconv.Atoi(payload["gas_limit"].(string))
            if err != nil {
                http.Error(w, "Failed to parse gas limit", http.StatusBadRequest)
                return
            }

            seed, err := strconv.Atoi(payload["random_id"].(string))
	    if err != nil {
                http.Error(w, "Bad Request, random_id", http.StatusBadRequest)
                return
            }

            //data := r.FormValue("data")

	    fmt.Printf("seed [%d]\n", seed)

	    rand.Seed(int64(seed))

	    // generate a random namespace ID
	    nID := generateRandHexEncodedNamespaceID()

	    // generate a random hex-encoded message
	    msg := generateRandMessage()

            fmt.Printf("http://%s:26659/submit_pfb %d %d %d %s", url, nID, gasLimit, fee, msg)
            responseBody, err := makeRequest(fmt.Sprintf("http://%s:26659/submit_pfb", url), nID, gasLimit, fee, msg)
            if err != nil {
                http.Error(w, fmt.Sprintf("Failed to make request: %s", err), http.StatusInternalServerError)
                return
            }

            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)

            w.Write(responseBody)

	    // OR
	    // mock
	    // read file out.json
	    // data, err := ioutil.ReadFile("out.json")
	    // if err != nil {
	    //     panic(err)
	    // }
	    // w.Write(data)

            return
        }

        tmpl := template.Must(template.ParseFiles("form.html"))
//	cess-Control-Allow-Origin'
	//req.Header.Set("Content-Type", "application/json")
        err := tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, fmt.Sprintf("Failed to render template: %s", err), http.StatusInternalServerError)
            return
        }
    })

    fmt.Println("Server listening on :8080")
    http.ListenAndServe(":8080", nil)
}

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

	// hack value with the namespaceId
	var jsonData map[string]interface{}
        err = json.Unmarshal(respBodyBytes, &jsonData)
        if err != nil {
            panic(err)
        }

        // Add key-value pair to the Go value
        jsonData["namespaceId"] = namespaceID

        // Marshal the updated Go value back to a JSON byte slice
        updatedJsonBytes, err := json.Marshal(jsonData)
        if err != nil {
            panic(err)
        }
	
	// Return response body
	//return respBodyBytes, nil
	return updatedJsonBytes, nil
}

