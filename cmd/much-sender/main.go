package main

import (
	"log"
	"net/http"

	"github.com/dogeorg/much-sender/internal/config"
	"github.com/dogeorg/much-sender/internal/email"
)

func main() {
	// Such Load configuration
	cfg, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatalf("Much Sad! Error loading config: %v", err)
	}

	// much Set up HTTP server
	http.HandleFunc("/send-email", email.Handler(cfg))

	log.Println("Much Wow! Starting much-sender server on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Much Sad! Server failed: %v", err)
	}
}
