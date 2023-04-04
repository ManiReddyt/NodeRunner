package NodeRunner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AppendEntries(term int, leaderID int, prevLogIndex int, prevLogTerm int, entries []LogEntry, leaderCommit int) {

	var data AppendEntriesArgs
	data.Term = term
	data.LeaderId = leaderID
	data.PrevLogIndex = prevLogIndex
	data.PrevLogTerm = prevLogTerm
	data.Entries = entries
	data.LeaderCommit = leaderCommit

	postBody, _ := json.Marshal(data)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://localhost:8080", "application/json", responseBody)
	if err != nil {
		log.Printf("response body is failed")
	}
	fmt.Println(resp)

}
