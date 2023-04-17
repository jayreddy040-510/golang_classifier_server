package main

import (
	"fmt"
	"net/http"
    "github.com/jayreddy040-510/golang_classifier_server/handlers"
)

    func main() {
        fmt.Println("starting up server...")
        http.HandleFunc("/detectspam", handlers.SmsSpamDetectionHandler)
        http.ListenAndServe(":7777", nil)
    }



