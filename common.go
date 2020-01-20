package gogithooks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// StatusCode is the value git recieves from the hook.
type StatusCode int

// Available Status codes
const (
	Success StatusCode = 0
	Abort   StatusCode = 1
)

// PushRef describes a single line of pre-receive input
type PushRef struct {
	OldValue string
	NewValue string
	RefName  string
}

func getPushOptions() []string {
	opts := []string{}

	if numOptions := os.Getenv("GIT_PUSH_OPTION_COUNT"); numOptions != "" {
		n, err := strconv.Atoi(numOptions)
		if err != nil || n < 0 {
			return opts
		}
		for i := 0; i < n; i++ {
			optName := fmt.Sprintf("GIT_PUSH_OPTION_%d", i)
			opts = append(opts, os.Getenv(optName))
		}
	}
	return opts
}

func getPushRefs() []PushRef {
	refs := []PushRef{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		refs = append(refs, PushRef{fields[0], fields[1], fields[2]})
	}
	return refs
}
