package gogithooks

import "os"

// PreMergeCommitArgs is the arguments given by git to the pre-merge-commit hook
type PreMergeCommitArgs struct {
	GitEditor string
}

// PreMergeCommit creates a hook for pre-merge-commit
func PreMergeCommit(handler func(args *PreMergeCommitArgs) StatusCode) {
	status := handler(&PreMergeCommitArgs{
		GitEditor: os.Getenv("GIT_EDITOR"),
	})
	os.Exit(int(status))
}
