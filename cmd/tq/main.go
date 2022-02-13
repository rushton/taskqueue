package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("print head of queue")
		return
	}

	cmd := args[0]

	if cmd == "done" {
		fmt.Println("Done: head of queue\nNext: next in queue")
		fmt.Println("Pop head")
		return
	}

	task := strings.Join(args, " ")
	fmt.Printf("taskqueue <- '%s'\n", task)
}
