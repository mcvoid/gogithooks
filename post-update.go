package gogithooks

import "os"

// PostUpdateArgs is the arguments given by git to the post-update hook
type PostUpdateArgs struct {
	UpdatedRefs []string
}

// PostUpdate creates a hook for post-update
func PostUpdate(handler func(args *PostUpdateArgs)) {
	args := PostUpdateArgs{
		UpdatedRefs: os.Args[1:],
	}
	handler(&args)
	os.Exit(0)
}
