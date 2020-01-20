package gogithooks

import (
	"log"
	"os"
)

// UpdateArgs are the arguments to the update handler
type UpdateArgs struct {
	RefName       string
	OldObjectName string
	NewObjectName string
}

// Update creates a handler for the update hook
func Update(handler func(args *UpdateArgs) StatusCode) {
	if len(os.Args) != 4 {
		log.Fatal("update: wrong number of command line args")
	}
	args := UpdateArgs{
		RefName:       os.Args[1],
		OldObjectName: os.Args[2],
		NewObjectName: os.Args[3],
	}

	status := handler(&args)
	os.Exit(int(status))
}
