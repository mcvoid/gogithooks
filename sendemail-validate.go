package gogithooks

import (
	"log"
	"os"
)

// SendemailValidateArgs is the arguments given by git to the sendemail-validate hook
type SendemailValidateArgs struct {
	EmailMessage *os.File
}

// SendemailValidate creates a hook for sendemail-validate
func SendemailValidate(handler func(args *SendemailValidateArgs) StatusCode) {
	if len(os.Args) < 2 {
		log.Fatal("sendemail-validate: not enough command line args")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("sendemail-validate: cannot open email message file")
	}
	defer file.Close()

	status := handler(&SendemailValidateArgs{
		EmailMessage: file,
	})
	os.Exit(int(status))
}
