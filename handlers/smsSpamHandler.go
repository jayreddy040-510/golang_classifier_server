package handlers

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
    func SmsSpamDetectionHandler(w http.ResponseWriter, r *http.Request) {
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
        cmd := exec.Command("python3", "sms_spam_classifier.py", sms)
        output, err := cmd.Output()

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        fmt.Printf("For SMS: \"%s\", \nSpam bool: %s", sms, string(output))
        w.Write(output)
    }


    func init() {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file", err)
        }
    }



