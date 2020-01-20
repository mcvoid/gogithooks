package gogithooks

import (
	"log"
	"os"
)

// PostMergeArgs is the arguments given by git to the post-merge hook
type PostMergeArgs struct {
	IsSquash bool
}

// PostMerge creates a hook for post-merge
func PostMerge(handler func(args *PostMergeArgs)) {
	if len(os.Args) != 2 {
		log.Fatal("post-merge: wrong number of command line args")
	}

	handler(&PostMergeArgs{
		IsSquash: os.Args[1] == "1",
	})
	os.Exit(0)
}
