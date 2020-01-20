package gogithooks

import (
	"log"
	"os"
)

// CommitMsgArgs is the arguments given by git to the commit-msg hook
type CommitMsgArgs struct {
	CommitMessage *os.File
}

// CommitMsg creates a hook for commit-msg
func CommitMsg(handler func(args *CommitMsgArgs) StatusCode) {
	if len(os.Args) < 2 {
		log.Fatal("commit-msg: not enough command line args")
	}

	args := CommitMsgArgs{}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("commit-msg: cannot open commit message file")
	}
	defer file.Close()
	args.CommitMessage = file

	status := handler(&args)
	os.Exit(int(status))
}
