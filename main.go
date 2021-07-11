package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Count struct {
	lines int
	words int
	size  int
}

func alphaCheck(r rune) int {
	switch {
	case (97 <= r && r <= 122) || (65 <= r && r <= 90) || (48 <= r && r <= 57):
		return 0
	case r == 32:
		return 1
	default:
		return 2
	}
}

func counter(reader bufio.Reader) Count {
	var count Count

	word := 0
	for {
		if r, size, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				if word != 0 {
					count.words++
				}
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if r == '\n' {
				count.lines++
				count.words++
				word = 0
			} else if alphaCheck(r) == 1 {
				count.words++
			} else if alphaCheck(r) == 0 {
				word++
			}

			count.size += size
		}
	}
	return count
}

func main() {

	optionLines := flag.Bool("l", false, "print the newline counts")
	optionBytes := flag.Bool("c", false, "print the byte counts")
	optionWords := flag.Bool("w", false, "print the word counts")
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

	reader := bufio.NewReader(file)

	count := counter(*reader)

	output := ""

	if *optionWords {
		output += fmt.Sprintf(" %d", count.words)
	}
	if *optionLines {
		output += fmt.Sprintf(" %d", count.lines)
	}
	if *optionBytes {
		output += fmt.Sprintf(" %d", count.size)
	}

	fmt.Printf("%s %s\n", output, fileName)
}
