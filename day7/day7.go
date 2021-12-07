package main

import (
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"math"
	"strings"
)

//go:embed input1
var input string

func main() {
	//input = "16,1,2,0,4,2,7,1,2,14\n"

	fields := strings.FieldsFunc(strings.TrimSpace(input), func(r rune) bool {
		return r == ','
	})

	positions := make([]int, len(fields))
	max := 0
	for i, f := range fields {
		p := aoc.ParseInt(f)
		positions[i] = p
		if p > max {
			max = p
		}
	}

	fmt.Printf("part one: %d\n", solve(positions, max, func(d int) int {
		return d
	}))
	fmt.Printf("part two: %d\n", solve(positions, max, func(d int) int {
		return (d + 1) * d / 2
	}))
}

func solve(positions []int, max int, dist func(int) int) int {
	min := math.MaxInt
	for x := 0; x <= max; x++ {
		fuel := 0
		for _, p := range positions {
			d := aoc.Abs(x - p)
			fuel += dist(d)
		}
		if fuel < min {
			min = fuel
		}
	}

	return min
}
