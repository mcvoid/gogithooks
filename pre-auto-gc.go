package gogithooks

import "os"

// PreAutoGc creates a hook for pre-auto-gc
func PreAutoGc(handler func() StatusCode) {
	status := handler()
	os.Exit(int(status))
}
