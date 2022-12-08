package main

import (
	"fmt"
	"github.com/hollesse/powerful-cli-book-exercises/chapter2/todo"
	"os"
	"strings"
)

const todoFileName = ".todo.json"

func main() {
	list := &todo.List{}
	if err := list.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *list {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		list.Add(item)
		if err := list.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
