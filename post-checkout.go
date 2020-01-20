package gogithooks

import (
	"log"
	"os"
)

// PostCheckoutArgs is the arguments given by git to the post-checkout hook
type PostCheckoutArgs struct {
	PreviousHead string
	NewHead      string
	IsBranch     bool
}

// PostCheckout creates a hook for post-checkout
func PostCheckout(handler func(args *PostCheckoutArgs)) {
	if len(os.Args) != 4 {
		log.Fatal("post-checkout: wrong number of command line args")
	}

	handler(&PostCheckoutArgs{
		PreviousHead: os.Args[1],
		NewHead:      os.Args[2],
		IsBranch:     os.Args[3] == "1",
	})
	os.Exit(0)
}
