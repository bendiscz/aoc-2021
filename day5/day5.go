package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"io"
	"regexp"
)

//go:embed input1
var input []byte

func main() {
	fmt.Printf("part one: %d\n", solve(bytes.NewReader(input), false))
	fmt.Printf("part two: %d\n", solve(bytes.NewReader(input), true))
}

var pattern = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func dir(a1, a2 int) int {
	if a1 < a2 {
		return 1
	} else {
		return -1
	}
}

func solve(reader io.Reader, pt2 bool) int {
	const D = 1000
	var screen [D][D]int

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		m := pattern.FindStringSubmatch(line)
		if m == nil {
			continue
		}

		x1, y1 := aoc.ParseInt(m[1]), aoc.ParseInt(m[2])
		x2, y2 := aoc.ParseInt(m[3]), aoc.ParseInt(m[4])

		if x1 == x2 {
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				screen[x1][y]++
			}
		} else if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				screen[x][y1]++
			}
		} else if pt2 && aoc.Abs(x2-x1) == aoc.Abs(y2-y1) {
			// part two
			dx := dir(x1, x2)
			dy := dir(y1, y2)
			for i := 0; i <= aoc.Abs(x2-x1); i++ {
				screen[x1+dx*i][y1+dy*i]++
			}
		}
	}

	count := 0
	for x := 0; x < D; x++ {
		for y := 0; y < D; y++ {
			if screen[x][y] > 1 {
				count++
			}
		}
	}

	return count
}
