package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func main() {
	SetupInterruptHandler()
	cfg := GetConfig()
	
	fmt.Printf("Satirical Seed Eater (%v)\n", cfg.Version)
	fmt.Println("--------------------------------")
	fmt.Println("Listening on port http://localhost:8080...")

	http.HandleFunc("/", handleRootRequest)
	http.ListenAndServe(":8080", nil)
}

type Introduction struct {
	Name string `json:"name"`
}

func handleRootRequest(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(res, "Error: %v", err)
		return
	}
	var intro Introduction
	err = json.Unmarshal(body, &intro)
	if err != nil {
		fmt.Fprintf(res, "Error: %v", err)
		return
	}
	fmt.Fprintf(res, "Hello, %s!\n", intro.Name)
}