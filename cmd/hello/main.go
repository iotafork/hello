package main

import (
	"fmt"
	"os"
	"uuid"

	"github.com/iotafork/hello/internal/greet"
)

func main() {
	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	msg, err := greet.Hello(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Printf("[%s] %s\n", uuid.New().String(), msg)
}
