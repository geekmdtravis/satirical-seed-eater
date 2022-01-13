package main

import (
	"fmt"
	"net/http"
)


func main() {
	cfg := GetConfig()
	
	fmt.Printf("Satirical Seed Eater (%v)\n", cfg.Version)
	fmt.Println("--------------------------------")
	fmt.Println("Listening on port http://localhost:8080...")

	http.HandleFunc("/", handleRootRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRootRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}