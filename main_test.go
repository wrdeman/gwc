package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLineCount(t *testing.T) {
	const input = "1\n2\n3\n"
	scanner := bufio.NewReader(strings.NewReader(input))

	count := counter(*scanner)
	assert.Equal(t, count.lines, 3, "should have 3 lines")
}

func TestByteCount(t *testing.T) {
	const input = "1\n2\n3\n"
	scanner := bufio.NewReader(strings.NewReader(input))

	count := counter(*scanner)
	assert.Equal(t, count.size, 6, "should have 6 bytes")

}

func TestWordCount(t *testing.T) {
	const input = "this should have\n5 words"
	scanner := bufio.NewReader(strings.NewReader(input))

	count := counter(*scanner)
	assert.Equal(t, count.words, 5, "should have 5 words")

}
