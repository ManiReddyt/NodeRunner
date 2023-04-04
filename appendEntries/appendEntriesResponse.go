package NodeRunner

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LogEntry struct {
	Term    int           `json:"term"`
	Command []interface{} `json:"command"`
}

type AppendEntriesArgs struct {
	Term         int        `json:"term"`
	LeaderId     int        `json:"leaderId"`
	PrevLogIndex int        `json:"prevLogIndex"`
	PrevLogTerm  int        `json:"prevLogTerm"`
	Entries      []LogEntry `json:"entries"`
	LeaderCommit int        `json:"leaderCommit"`
}

type AppendEntriesResult struct {
	Term    int  `json:"term"`
	Success bool `json:"success"`
}

func AppendEntriesResponse(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	var data AppendEntriesArgs

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}

	var response AppendEntriesResult

	err = json.NewEncoder(writer).Encode(&response)
	fmt.Println(data)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}

}

// {
//     "Term":1,
//     "LeaderID":0,
//     "PrevLogIndex":0,
//     "PrevLogTerm":0,
//     "Entries":{},
//     "LeaderCommit":0
// }
