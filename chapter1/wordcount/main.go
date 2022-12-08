package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count Bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(reader io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(reader)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wordcount := 0
	for scanner.Scan() {
		wordcount++
	}

	return wordcount
}
