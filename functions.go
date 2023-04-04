package NodeRunner

func BuildRequestVote(rf *raft) RequestVoteArgs {
	response := RequestVoteArgs{
		Term:         rf.PersistentServerState.CurrentTerm,
		CandidateId:  rf.me,
		LastLogIndex: rf.VolatileServerState.CommitIndex,
		LastLogTerm:  rf.PersistentServerState.CurrentTerm,
	}
	return response
}

func (rf *raft) HandleRequestVote(requestVote *RequestVoteArgs) RequestVoteResult {
	response := RequestVoteResult{}
	if requestVote.Term < rf.PersistentServerState.CurrentTerm {
		response.VoteGranted = false
	}
	if requestVote.Term > rf.PersistentServerState.CurrentTerm {
		rf.PersistentServerState.CurrentTerm = requestVote.Term
		rf.PersistentServerState.VotedFor = -1
	}
	if rf.PersistentServerState.VotedFor == -1 || rf.PersistentServerState.VotedFor == requestVote.CandidateId {
		lastLogIndex := len(rf.PersistentServerState.Log) - 1
		lastLogTerm := -1
		if lastLogIndex >= 0 {
			lastLogTerm = rf.PersistentServerState.Log[lastLogIndex].Term
		}

		if requestVote.LastLogIndex > lastLogTerm && (requestVote.LastLogTerm == lastLogTerm && requestVote.LastLogIndex >= lastLogIndex) {
			response.VoteGranted = true
			rf.PersistentServerState.VotedFor = requestVote.CandidateId
		}
	}
	response.Term = rf.PersistentServerState.CurrentTerm
	return response
}

func (rf *raft) HandleRequestVoteResult(requestVoteResult *RequestVoteResult) {

	// if rf.VolatileServerState.CommitIndex > rf.VolatileServerState.LastApplied {
	// 	rf.VolatileServerState.LastApplied += 1
	// 	//apply log[lastApplied] to state machine
	// }
	if requestVoteResult.Term > rf.PersistentServerState.CurrentTerm {
		rf.PersistentServerState.CurrentTerm = requestVoteResult.Term
		rf.status = "Follower"
	}
	if requestVoteResult.VoteGranted {
		rf.count += 1
		if rf.count >= ((rf.totalNodes) / 2) {
			rf.status = "Leader"
		}
	}

	// if rf.status == "Follower" {

	// }
	// if rf.status == "Candidate" {
	// 	rf.PersistentServerState.CurrentTerm += 1

	// }
}

func (rf *raft) BuildAppendEntries() AppendEntriesArgs {
	response := AppendEntriesArgs{
		Term:         rf.PersistentServerState.CurrentTerm,
		LeaderId:     rf.me,
		PrevLogIndex: len(rf.PersistentServerState.Log) - 1,
		PrevLogTerm:  rf.PersistentServerState.Log[len(rf.PersistentServerState.Log)-1].Term,
		Entries:      rf.PersistentServerState.Log[len(rf.PersistentServerState.Log)-1].Command,
		LeaderCommit: rf.VolatileServerState.CommitIndex,
	}
	return response
}

func (rf *raft) HandleAppendEntries(AppendEntries *AppendEntriesArgs) AppendEntriesResult {
	var response AppendEntriesResult
	if rf.PersistentServerState.CurrentTerm > AppendEntries.Term {
		response.Success = false
	}
	if ((len(rf.PersistentServerState.Log) - 1) < AppendEntries.PrevLogIndex) && (rf.PersistentServerState.CurrentTerm == AppendEntries.Term) {
		response.Success = false
	}
	if (len(rf.PersistentServerState.Log) - 1) == AppendEntries.PrevLogIndex+1 {
		rf.PersistentServerState.Log[AppendEntries.PrevLogIndex+1].Command = AppendEntries.Entries
	}

	//TODO:append any new entries not already in th log

	if AppendEntries.LeaderCommit > rf.VolatileServerState.CommitIndex {
		if AppendEntries.LeaderCommit < rf.VolatileServerState.CommitIndex {
			rf.VolatileServerState.CommitIndex = AppendEntries.LeaderCommit
		} else {
			rf.VolatileServerState.CommitIndex = len(rf.PersistentServerState.Log) - 1
		}
	}
	response.Term = rf.PersistentServerState.CurrentTerm
	return response
}

func (rf *raft) HandleAppendEntriesResult(AppendResult *AppendEntriesResult) {
	if AppendResult.Term > rf.PersistentServerState.CurrentTerm {
		rf.PersistentServerState.CurrentTerm = AppendResult.Term
		rf.status = "Follower"
	}
	if AppendResult.Success {
	}
}
