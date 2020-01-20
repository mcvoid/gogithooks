package gogithooks

import (
	"log"
	"os"
)

// PreRebaseArgs is the arguments given by git to the pre-rebase hook
type PreRebaseArgs struct {
	Upstream           string
	BranchBeingRebased string
}

// PreRebase creates a hook for pre-rebase
func PreRebase(handler func(args *PreRebaseArgs) StatusCode) {
	if len(os.Args) < 2 {
		log.Fatal("pre-rebase: not enough command line args")
	}

	args := PreRebaseArgs{
		Upstream: os.Args[1],
	}

	if len(os.Args) > 2 {
		args.BranchBeingRebased = os.Args[2]
	}

	status := handler(&args)
	os.Exit(int(status))
}
