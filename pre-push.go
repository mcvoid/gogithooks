package gogithooks

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// PrePushRef describes a single line of pre-push input
type PrePushRef struct {
	LocalRef   string
	LocalSha1  string
	RemoteRef  string
	RemoteSha1 string
}

// PrePushArgs is the arguments given by git to the pre-push hook
type PrePushArgs struct {
	RemoteName   string
	RemoteSource string
	Objects      []PrePushRef
}

// PrePush creates a hook for pre-push
func PrePush(handler func(args *PrePushArgs) StatusCode) {
	if len(os.Args) != 3 {
		log.Fatal("pre-push: wrong number of command line args")
	}
	args := PrePushArgs{}
	args.RemoteName = os.Args[1]
	args.RemoteSource = os.Args[2]

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		args.Objects = append(args.Objects, PrePushRef{fields[0], fields[1], fields[2], fields[3]})
	}

	status := handler(&args)
	os.Exit(int(status))
}
