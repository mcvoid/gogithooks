package gogithooks

import (
	"log"
	"os"
)

// PushToCheckoutArgs is the arguments given by git to the push-to-checkout hook
type PushToCheckoutArgs struct {
	CurrentCommit string
}

// PushToCheckout creates a hook for push-to-checkout
func PushToCheckout(handler func(args *PushToCheckoutArgs) StatusCode) {
	if len(os.Args) != 2 {
		log.Fatal("push-to-checkout: wrong number of command line args")
	}
	args := PushToCheckoutArgs{
		CurrentCommit: os.Args[1],
	}
	status := handler(&args)
	os.Exit(int(status))
}
