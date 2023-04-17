package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
    "io"
    "encoding/json"
    "log"
    "github.com/joho/godotenv"
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
            fmt.Println(os.Getenv("API_KEY"), apiKey)
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
        // fmt.Println(sms)
        cmd := exec.Command("python3", "sms_spam_classifier.py", sms)
        output, err := cmd.CombinedOutput()

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        // fmt.Println("model output", output)
        w.Write(output)
    }


    func init() {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file", err)
        }
    }


    func main() {
        fmt.Println("starting up server...")
        http.HandleFunc("/detectspam", smsSpamDetectionHandler)
        http.ListenAndServe(":7777", nil)
    }



