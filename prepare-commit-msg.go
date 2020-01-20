package gogithooks

import (
	"log"
	"os"
)

// CommitMessageSource represents the source of a commit message
type CommitMessageSource string

// Available CommitSources
const (
	MessageSource  CommitMessageSource = "message"
	TemplateSource CommitMessageSource = "template"
	MergeSource    CommitMessageSource = "merge"
	SquashSource   CommitMessageSource = "squash"
	CommitSource   CommitMessageSource = "commit"
)

// PrepareCommitMsgArgs is the arguments given by git to the prepare-commit-msg hook
type PrepareCommitMsgArgs struct {
	CommitMessage *os.File
	Source        CommitMessageSource
	CommitID      string
}

// PrepareCommitMsg creates a hook for prepare-commit-msg
func PrepareCommitMsg(handler func(args *PrepareCommitMsgArgs) StatusCode) {
	if len(os.Args) < 2 {
		log.Fatal("prepare-commit-msg: not enough command line args")
	}

	args := PrepareCommitMsgArgs{}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("prepare-commit-msg: cannot open commit message file")
	}
	defer file.Close()
	args.CommitMessage = file

	if len(os.Args) > 2 {
		args.Source = CommitMessageSource(os.Args[2])
	}

	if len(os.Args) > 3 {
		args.CommitID = os.Args[3]
	}

	status := handler(&args)
	os.Exit(int(status))
}
