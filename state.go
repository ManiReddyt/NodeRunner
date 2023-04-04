package NodeRunner

import "time"

// var persistentState PersistentState
// var volatileState VolatileState
// var volatileStateOnLeaders VolatileStateonLeaders

type AppendEntriesArgs struct {
	Term         int           `json:"term"`
	LeaderId     int           `json:"leaderId"`
	PrevLogIndex int           `json:"prevLogIndex"`
	PrevLogTerm  int           `json:"prevLogTerm"`
	Entries      []interface{} `json:"entries"`
	LeaderCommit int           `json:"leaderCommit"`
}

type AppendEntriesResult struct {
	Term    int  `json:"term"`
	Success bool `json:"success"`
}
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

type LogEntry struct {
	Term    int           `json:"term"`
	Command []interface{} `json:"command"`
}

type PersistentServerState struct {
	CurrentTerm int        `json:"currentTerm"`
	VotedFor    int        `json:"votedFor"`
	Log         []LogEntry `json:"log"`
}

type VolatileServerState struct {
	CommitIndex int `json:"commitIndex"`
	LastApplied int `json:"lastApplied"`
}

type VolatileStateonLeaders struct {
	NextIndex  []int `json:"nextIndex"`
	MatchIndex []int `json:"matchIndex"`
}

type raft struct {
	PersistentServerState
	VolatileServerState
	VolatileStateonLeaders
	peers        []int
	status       string
	lastAccessed time.Time
	me           int
	lastLogIndex int
	lastLogTerm  int
	count        int
	totalNodes   int
}

func starter(me int, peers []string) *raft {
	rf := &raft{}
	rf.status = "Follower"
	rf.Log = []LogEntry{
		{
			Term:    0,
			Command: nil,
		},
	}
	rf.NextIndex = make([]int, 5)
	rf.MatchIndex = make([]int, 5)
	rf.VotedFor = -1
	rf.me = me

	go rf.manageLifecycle()
	return rf
}

func (rf *raft) manageLifecycle() {
	for true {
		status := rf.status
		if status == "Follower" {
			rf.manageFollower()
		} else if status == "Candidate" {
			rf.manageCandidate()
		} else {
			rf.manageLeader()
		}
	}
}
