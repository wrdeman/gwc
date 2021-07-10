package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLineCount(t *testing.T) {
	const input = "1\n2\n3\n"
	scanner := bufio.NewScanner(strings.NewReader(input))

	count := counter(*scanner, true)
	assert.Equal(t, count.lines, 3, "should have 3 lines")
}
