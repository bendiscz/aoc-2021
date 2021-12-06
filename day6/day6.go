package main

import (
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"unicode"
)

//go:embed input1
var input string

var dayCount = simFish(256)

func simFish(days int) []int {
	day := make([]int, days+1)
	fish := []int{0, 1, 0, 0, 0, 0, 0, 0, 0}

	for d := 0; d <= days; d++ {
		born := fish[0]
		count := 2 * born
		for i := 1; i < len(fish); i++ {
			count += fish[i]
			fish[i-1] = fish[i]
		}
		fish[6] += born
		fish[8] = born
		day[d] = count
	}

	return day
}

func main() {
	count80, count256 := 0, 0
	for _, c := range input {
		if unicode.IsDigit(c) {
			count80 += dayCount[80-aoc.ParseInt(string(c))]
			count256 += dayCount[256-aoc.ParseInt(string(c))]
		}
	}
	fmt.Printf("part one: %d\n", count80)
	fmt.Printf("part two: %d\n", count256)
}
