package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type Count struct {
	lines int
	words int
}

func counter(scanner bufio.Scanner, optionLine bool) Count {
	var count Count

	lines := 0
	for scanner.Scan() {
		if optionLine {
			lines++
		}
	}
	count.lines = lines
	return count
}

func main() {

	optionLine := flag.Bool("l", false, "count lines")
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("Please provide path to file")
	} else if flag.NArg() > 1 {
		log.Fatal("Please provide path to one file")
	}
	fileName := flag.Arg(0)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	count := counter(*scanner, *optionLine)

	fmt.Printf("%d %s\n", count.lines, fileName)
}
