package NodeRunner

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestVoteArgs struct {
	Term         int `json:"term"`
	CandidateId  int `json:"candidateId"`
	LastLogIndex int `json:"lastLogIndex"`
	LastLogTerm  int `json:"lastLogTerm"`
}

type RequestVoteResult struct {
	Term        int  `json:"term"`
	VoteGranted bool `json:"voteGranted"`
}

func requestVoteResponse(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	var data RequestVoteArgs
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}

	var dummy RequestVoteResult
	dummy.Term = 5
	dummy.VoteGranted = true

	err = json.NewEncoder(writer).Encode(&dummy)
	fmt.Println(data)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct request vote")
	}
}
