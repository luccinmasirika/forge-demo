package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "v1"

func main() {
	started := time.Now().UTC().Format(time.RFC3339)
	host, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Forge!\nversion: %s\nstarted: %s\ncontainer: %s\n", version, started, host)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	log.Printf("forge-demo %s listening on :8080", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
