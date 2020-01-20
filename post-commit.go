package gogithooks

import "os"

// PostCommit creates a hook for post-applypatch
func PostCommit(handler func()) {
	handler()
	os.Exit(0)
}
