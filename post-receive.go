package gogithooks

import (
	"os"
)

// PostReceiveArgs is the arguments given by git to the pre-receive hook
type PostReceiveArgs struct {
	Refs    []PushRef
	Options []string
}

// PostReceive creates a hook for pre-receive
func PostReceive(handler func(args *PostReceiveArgs)) {
	args := PostReceiveArgs{
		Refs:    getPushRefs(),
		Options: getPushOptions(),
	}

	handler(&args)
	os.Exit(0)
}
