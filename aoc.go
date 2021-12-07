package aoc

import (
	"strconv"
)

func ParseInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
