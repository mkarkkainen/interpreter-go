package main

import (
    "fmt"
    "os"
    "os/user"
    "github.com/mkarkkainen/interpreter-go/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Howdy %s!\n", user.Username)
    fmt.Printf("Feel free to type in some commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
