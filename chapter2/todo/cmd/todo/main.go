package main

import (
	"flag"
	"fmt"
	"github.com/hollesse/powerful-cli-book-exercises/chapter2/todo"
	"os"
)

const todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed by Joshua Töpfer\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Copyright 2022")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	taskFlag := flag.String("task", "", "Task to be included in the ToDo list")
	listFlag := flag.Bool("list", false, "List all Tasks")
	completeFlag := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()
	list := &todo.List{}
	if err := list.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *listFlag:
		fmt.Print(list)
	case *completeFlag > 0:
		if err := list.Complete(*completeFlag); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := list.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *taskFlag != "":
		list.Add(*taskFlag)
		if err := list.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
