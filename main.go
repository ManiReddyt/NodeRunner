package NodeRunner

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/appendEntries", AppendEntriesResponse).Methods("POST")
	router.HandleFunc("/requestVote", requestVoteResponse).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
	// AppendEntries(0, 0, 0, 0, _, 0)

}
