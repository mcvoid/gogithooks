package gogithooks

import (
	"os"
)

// PreReceiveArgs is the arguments given by git to the pre-receive hook
type PreReceiveArgs struct {
	Refs    []PushRef
	Options []string
}

// PreReceive creates a hook for pre-receive
func PreReceive(handler func(args *PreReceiveArgs) StatusCode) {
	args := PreReceiveArgs{
		Refs:    getPushRefs(),
		Options: getPushOptions(),
	}

	status := handler(&args)
	os.Exit(int(status))
}
