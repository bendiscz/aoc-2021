package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"regexp"
)

var linePattern = regexp.MustCompile(`^(forward|up|down) (\d+)$`)

//go:embed input1
var input []byte

func main() {
	fmt.Printf("part one: %v\n", solve(bufio.NewScanner(bytes.NewReader(input)), false))
	fmt.Printf("part one: %v\n", solve(bufio.NewScanner(bytes.NewReader(input)), true))
}

func solve(scanner *bufio.Scanner, aiming bool) int {
	forward, depth, aim := 0, 0, 0

	for scanner.Scan() {
		m := linePattern.FindStringSubmatch(scanner.Text())
		if m == nil {
			continue
		}
		val := aoc.ParseInt(m[2])
		switch m[1] {
		case "forward":
			forward += val
			depth += aim * val
		case "down":
			if aiming {
				aim += val
			} else {
				depth += val
			}
		case "up":
			if aiming {
				aim -= val
			} else {
				depth -= val
			}
		}
	}

	return forward * depth
}
