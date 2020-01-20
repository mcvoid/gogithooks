package gogithooks

import (
	"os"
)

// P4PreSubmit creates a hook for p4-pre-submit
func P4PreSubmit(handler func() StatusCode) {
	status := handler()
	os.Exit(int(status))
}
