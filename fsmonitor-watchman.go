package gogithooks

import (
	"log"
	"os"
	"strconv"
	"time"
)

// FsmonitorWatchmanArgs is the arguments given by git to the fsmonitor-watchman hook
type FsmonitorWatchmanArgs struct {
	Version string
	Time    time.Time
}

// FsmonitorWatchman creates a hook for fsmonitor-watchman
func FsmonitorWatchman(handler func(args *FsmonitorWatchmanArgs) StatusCode) {
	if len(os.Args) != 3 {
		log.Fatal("fsmonitor-watchman: wrong number of command line args")
	}

	timestamp, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal("fsmonitor-watchman: invalid timestamp")
	}

	status := handler(&FsmonitorWatchmanArgs{
		Version: os.Args[1],
		Time:    time.Unix(0, timestamp),
	})
	os.Exit(int(status))
}
