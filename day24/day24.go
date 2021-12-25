package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

//go:embed input0
var input []byte

var pattern = regexp.MustCompile(`^$`)

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		m := pattern.FindStringSubmatch(scanner.Text())
		_, _ = line, m
	}

	fields := strings.FieldsFunc(string(input), func(r rune) bool {
		return unicode.IsSpace(r) || r == ','
	})
	for _, f := range fields {
		_ = f
	}

	fmt.Printf("part one: %v\n", 0)
	fmt.Printf("part two: %v\n", 0)
}
