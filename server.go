package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
    "io"
    "encoding/json"
)
    type RequestBody struct {
        SMS string `json:"sms"`
    }
    func smsSpamDetectionHandler(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
            return
        }

        apiKey := r.Header.Get("apiKey")
        if apiKey != os.Getenv("API_KEY") {
            http.Error(w, "You are not authorized to access this resource.", http.StatusForbidden)
            return
        }

        body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
        defer r.Body.Close()
        if err != nil {
            http.Error(w, "Error reading request body.", http.StatusInternalServerError)
            return
        }
        var requestBody RequestBody
        err = json.Unmarshal(body, &requestBody)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        sms := requestBody.SMS
        fmt.Fprintf(w, "SMS receieved: %s", sms)


        cmd := exec.Command("python3", "sms_spam_classifier.py", sms)
        output, err := cmd.Output()

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        w.Write(output)
    }

    func main() {
        fmt.Println("starting up server...")
        http.HandleFunc("/detectspam", smsSpamDetectionHandler)
        http.ListenAndServe(":7777", nil)
    }



