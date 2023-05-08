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
        if r.Method == "POST" {
            err := r.ParseForm()
            if err != nil {
                http.Error(w, "Failed to parse form", http.StatusBadRequest)
                return
            }

            url := r.FormValue("url")
            seedStr := r.FormValue("random_id")
            gasLimitStr := r.FormValue("gas_limit")
            feeStr := r.FormValue("fee")
            //data := r.FormValue("data")

	    seed, err := strconv.Atoi(seedStr)
            if err != nil {
                http.Error(w, "Failed to parse randid", http.StatusBadRequest)
                return
            }
	    fmt.Printf("seed [%d]\n", seed)

	    rand.Seed(int64(seed))

	    // generate a random namespace ID
	    nID := generateRandHexEncodedNamespaceID()

	    // generate a random hex-encoded message
	    msg := generateRandMessage()

            gasLimit, err := strconv.Atoi(gasLimitStr)
            if err != nil {
                http.Error(w, "Failed to parse gas limit", http.StatusBadRequest)
                return
            }

            fee, err := strconv.Atoi(feeStr)
            if err != nil {
                http.Error(w, "Failed to parse fee", http.StatusBadRequest)
                return
            }

            // err = makeRequest(url, namespaceID, gasLimit, fee, data)
            responseBody, err := makeRequest(fmt.Sprintf("http://%s:26659/submit_pfb", url), nID, gasLimit, fee, msg)
            if err != nil {
                http.Error(w, fmt.Sprintf("Failed to make request: %s", err), http.StatusInternalServerError)
                return
            }

            // response := Response{Message: "Request successful"}
            // responseJSON, err := json.Marshal(response)
            // if err != nil {
            //     http.Error(w, fmt.Sprintf("Failed to marshal response: %s", err), http.StatusInternalServerError)
            //     return
            // }

            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            w.Write(responseBody)
            return
        }

        tmpl := template.Must(template.ParseFiles("form.html"))
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

	// Return response body
	return respBodyBytes, nil
}

