package main

import (
	"fmt"
	"net/http"

	"github.com/geekmdtravis/satirical-seed-eater/handlers"
	"github.com/geekmdtravis/satirical-seed-eater/utils"
)


func main() {

	SetupInterruptHandler()
	printWelcomeMessage()
	startServer()
}

func startServer() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/validate", handlers.HandleSeedValidationRequest)
	http.ListenAndServe(":8080", nil)
}

func printWelcomeMessage() {
	cfg := utils.GetConfig()
	fmt.Printf("Satirical Seed Eater (%v)\n", cfg.Version)
	fmt.Println("--------------------------------")
	fmt.Println("Listening on port http://localhost:8080...")
}
