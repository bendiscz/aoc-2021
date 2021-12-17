package main

import (
	_ "embed"
	"fmt"
)

const x1, x2 = 175, 227
const y1, y2 = -134, -79

func main() {
	mdy := -y1 - 1
	y := (1 + mdy) * mdy / 2
	fmt.Printf("part one: %v\n", y)

	count := 0
	for dx := x2; dx > 0; dx-- {
		for dy := mdy; dy >= y1; dy-- {
			if simulate(dx, dy) {
				count++
			}
		}
	}

	fmt.Printf("part two: %v\n", count)
}

func simulate(dx, dy int) bool {
	x, y := 0, 0
	for x <= x2 && y >= y1 {
		x += dx
		y += dy

		if x >= x1 && x <= x2 && y <= y2 && y >= y1 {
			return true
		}

		if dx > 0 {
			dx--
		}
		dy--
	}
	return false
}
