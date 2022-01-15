package main

import (
	"fmt"
	"net/http"
)


func main() {

	SetupInterruptHandler()
	printWelcomeMessage()
	startServer()
}

func startServer() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/validate", handleSeedValidationRequest)
	http.ListenAndServe(":8080", nil)
}

func printWelcomeMessage() {
	cfg := GetConfig()
	fmt.Printf("Satirical Seed Eater (%v)\n", cfg.Version)
	fmt.Println("--------------------------------")
	fmt.Println("Listening on port http://localhost:8080...")
}
