package NodeRunner

import (
	"math/rand"
	"time"
)

const Mintime = 1500
const Maxtime = 3000

func randTime() time.Duration {
	randTime := Mintime + rand.Intn(Maxtime-Mintime)
	return time.Duration(randTime) + time.Millisecond
}

func (rf *raft) manageFollower() {
	duration := randTime()
	time.Sleep(duration)
	lastAccessed := rf.lastAccessed

	if time.Now().Sub(lastAccessed).Milliseconds() >= duration.Milliseconds() {
		rf.status = "Candidate"
		rf.VotedFor = -1
		rf.CurrentTerm++
	}

}

func (rf *raft) manageCandidate() {
	duration := randTime()
	me := rf.me
	peers := rf.peers
	term := rf.CurrentTerm
	lastLongIndex := rf.lastLogIndex
	lastLogTerm := rf.lastLogTerm
	total := len(peers)
	count := 0
	finished := 0
	majority := (total / 2) + 1
	for peer := range peers {
		if me == peer {
			count++
			finished++
			continue
		}
		go func(peer int) {
			var args RequestVoteArgs
			args.Term = term
			args.CandidateId = me
			args.LastLogIndex = lastLogTerm
			args.LastLogTerm = lastLogTerm
			var reply RequestVoteArgs
			ok := rf.sendRequestVote(peer, &args, &reply)
			if !ok {
				finished++
				return
			}
			if reply.VoteGranted {
				finished++
				count++
			} else {
				finished++
				if args.Term < reply.Term {
					rf.status = "Follower"
				}
			}
		}(peer)
	}

}

func (rf *raft) manageLeader() {

}
