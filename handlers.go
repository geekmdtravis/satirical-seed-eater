package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type SeedValidationRequest struct {
	MagicWord string `json:"magicWord"`
	SeedPhrase string `json:"seedPhrase"`
}

type SeedValidationReply struct {
	Valid bool `json:"valid"`
	Comment string `json:"comment"`
}

func handleSeedValidationRequest(res http.ResponseWriter, req *http.Request) {
	reply := SeedValidationReply{Valid: false, Comment: ""}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		reply.Comment = "Error: %v" + err.Error()
		replyJson, _ := json.Marshal(reply)
		res.Write(replyJson)
		return
	}

	var intro SeedValidationRequest
	err = json.Unmarshal(body, &intro)
	if err != nil {
		reply.Comment = "Error: " + err.Error()
		replyJson, _ := json.Marshal(reply)
		fmt.Fprint(res, string(replyJson))
		return
	}
	if strings.ToLower(intro.MagicWord) != "please" {
		reply.Comment = "Error: Invalid magic word"
		replyJson, _ := json.Marshal(reply)
		fmt.Fprint(res, string(replyJson))
		return
	}

	seedWordCnt := len(strings.Split(intro.SeedPhrase, " "))
	if seedWordCnt != 12 && seedWordCnt != 24 {
		reply.Comment = "Error: Invalid seed phrase; should be either 12 or 24 words"
		replyJson, _ := json.Marshal(reply)
		fmt.Fprint(res, string(replyJson))
		return
	}

	if err != nil {
		reply.Comment = "Error: " + err.Error()
		replyJson, _ := json.Marshal(reply)
		fmt.Fprint(res, string(replyJson))
		return
	}
	
	reply.Comment = "Your seed phrase is valid, and it's safe with me! ;)"
	reply.Valid = true
	replyJson, _ := json.Marshal(reply)
	fmt.Fprint(res, string(replyJson))
}