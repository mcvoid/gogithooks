package gogithooks

import "os"

// PostApplypatch creates a hook for post-applypatch
func PostApplypatch(handler func()) {
	handler()
	os.Exit(0)
}
