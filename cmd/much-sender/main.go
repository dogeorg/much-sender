package main

import (
	"fmt"
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

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Much Wow! Starting much-sender server on %s", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Much Sad! Server failed: %v", err)
	}
}
