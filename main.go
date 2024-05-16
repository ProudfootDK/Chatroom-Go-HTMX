package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func main() {
	fmt.Println("hello world")

	hub := NewHub()
	go hub.Run()

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(hub, w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified
	}
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
