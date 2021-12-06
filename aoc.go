package aoc

import (
	"strconv"
)

func ParseInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
