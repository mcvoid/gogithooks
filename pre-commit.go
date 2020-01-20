package gogithooks

import "os"

// PreCommitArgs is the arguments given by git to the pre-commit hook
type PreCommitArgs struct {
	GitEditor string
}

// PreCommit creates a hook for pre-commit
func PreCommit(handler func(args *PreCommitArgs) StatusCode) {
	status := handler(&PreCommitArgs{
		GitEditor: os.Getenv("GIT_EDITOR"),
	})
	os.Exit(int(status))
}
