package NodeRunner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func requestVote(term int, candidateId int, lastLogIndex int, lastLogTerm int) {
	var data RequestVoteArgs
	data.Term = term
	data.CandidateId = candidateId
	data.LastLogIndex = lastLogIndex
	data.LastLogTerm = lastLogTerm

	postBody, _ := json.Marshal(data)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://localhost:8080", "application/json", responseBody)
	if err != nil {
		log.Printf("response body is failed")
	}
	fmt.Println(resp)
}
