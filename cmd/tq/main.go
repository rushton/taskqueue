package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rushton/taskqueue/pkg/tasks"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		head := tasks.Head()
		if head == nil {
			fmt.Fprintln(os.Stderr, "no tasks in queue")
			return
		}
		fmt.Println(head)
		return
	}

	cmd := args[0]

	if cmd == "done" {
		itm, err := tasks.Done()
		if err != nil {
			panic(err)
		}
		if itm == nil {
			fmt.Fprintln(os.Stderr, "no tasks in queue, nothing to complete.")
		} else {
			fmt.Println(itm)
		}
		return
	}
	if cmd == "list" {
		itms, err := tasks.List()
		if err != nil {
			panic(err)
		}
		for _, itm := range itms {
			fmt.Println(itm)
		}
		return
	}

	task := tasks.Item{
		Created:     time.Now(),
		Description: strings.Join(args, " "),
	}
	tasks.Put(task)
	fmt.Printf("taskqueue <- '%s'\n", task)
}
