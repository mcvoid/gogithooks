package gogithooks

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// RewrittenCommit is the info of a commit that has been rewritten.
type RewrittenCommit struct {
	OldSha1   string
	NewSha1   string
	ExtraInfo string
}

// PostRewriteArgs is the arguments given by git to the post-rewrite hook
type PostRewriteArgs struct {
	Command     string
	CommandArgs []string
	Commits     []RewrittenCommit
}

// PostRewrite creates a hook for post-rewrite
func PostRewrite(handler func(args *PostRewriteArgs)) {
	if len(os.Args) < 2 {
		log.Fatal("post-rewrite: wrong number of command line args")
	}
	args := PostRewriteArgs{}

	args.Command = os.Args[1]
	args.CommandArgs = os.Args[2:]

	commits := []RewrittenCommit{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) == 2 {
			commits = append(commits, RewrittenCommit{
				OldSha1: fields[0],
				NewSha1: fields[1],
			})
			continue
		}
		commits = append(commits, RewrittenCommit{
			OldSha1:   fields[0],
			NewSha1:   fields[1],
			ExtraInfo: fields[2],
		})
	}
	args.Commits = commits

	handler(&args)
	os.Exit(0)
}
