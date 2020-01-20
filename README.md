# Go Git Hooks

Go Git Hooks is a simple idiomatic API to implement git hooks.

Installation
------------
Install using ```go get``` (no dependecies required):
```
go get github.com/mcvoid/gogithooks
```

How to use
----------
1. Give the library a handler that you define.
2. Build the binary.
3. Rename the binary to the appropriate hook name.
4. Place the binary into ```.git/hooks```

Example
-------
```
package main

import (
  "fmt"

  hooks "github.com/mcvoid/gogithooks"
)

func PreRebaseHandler(args *hooks.PreRebaseArgs) hooks.StatusCode {
  fmt.Println("Running the pre-rebase hook")
  fmt.Println("Upstream:", args.Upstream)
  fmt.Println("Branch being rebased", args.BranchBeingRebased)

  return hooks.Success
}

func main() {
  hooks.PreRebase(PreRebaseHandler)
}
```

"Why would I want to use this library?"
---------------------------------------
Git provides information for hooks through several different sources: environment variables, command line arguments, and standard input. In some cases, it provides streams of text that need to be parsed, or file names that need to be opened. This library does that for you. It gathers your command line arguments, relevant environment variables, and input streams into a few function arguments. Then you don't have to constantly reference the native git hook API.

"How do I interact with git inside the handler?"
------------------------------------------------
This libray isn't trying to be a complete interface between your hook and git. It's only trying to provide a Go version of the hooks API git already provides. Instead you'e encouraged to bring your own solution to further interact with git's plumbing. You have a few options. For simple tasks, invloking git through the command line would be adequate. A better option would be to use [go-git](https://github.com/src-d/go-git).

License
-------
MIT License, see [LICENSE](LICENSE)