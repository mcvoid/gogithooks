package gogithooks

import (
	"log"
	"os"
)

// PostIndexChangeArgs is the arguments given by git to the post-index-change hook
type PostIndexChangeArgs struct {
	IsWorkingDirectoryUpdated bool
	IsSkipWorktreeBitsUpdated bool
}

// PostIndexChange creates a hook for post-index-change
func PostIndexChange(handler func(args *PostIndexChangeArgs) StatusCode) {
	if len(os.Args) != 3 {
		log.Fatal("post-index-change: incorrect number of command line args")
	}

	args := PostIndexChangeArgs{
		os.Args[1] == "1",
		os.Args[2] == "1",
	}
	if args.IsWorkingDirectoryUpdated && args.IsSkipWorktreeBitsUpdated {
		log.Fatal("post-index-change: both args whould never be true at the same time")
	}

	status := handler(&args)
	os.Exit(int(status))
}
