package gogithooks

import "os"

// PreApplypatch creates a hook for pre-applypatch
func PreApplypatch(handler func() StatusCode) {
	status := handler()
	os.Exit(int(status))
}
