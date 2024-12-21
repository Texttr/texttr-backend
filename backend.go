package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, "Error in body!")
		}
		bodyStr := string(body)
		fmt.Fprint(w, bodyStr)
		log.Println(bodyStr)
		defer r.Body.Close()
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/text", requestHandler)

	port := ":8080"
	fmt.Printf("[BACKEND] HTTP API is now listening on port %x", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
