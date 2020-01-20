package gogithooks

import (
	"log"
	"os"
)

// ApplypatchMsgArgs is the arguments given by git to the applypatch-msg hook
type ApplypatchMsgArgs struct {
	CommitMessage *os.File
}

// ApplypatchMsg creates a hook for applypatch-msg
func ApplypatchMsg(handler func(args *ApplypatchMsgArgs) StatusCode) {
	if len(os.Args) < 2 {
		log.Fatal("applypatch-msg: not enough command line args")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("applypatch-msg: cannot open commit message file")
	}
	defer file.Close()

	status := handler(&ApplypatchMsgArgs{
		CommitMessage: file,
	})
	os.Exit(int(status))
}
