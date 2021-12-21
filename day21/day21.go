package main

import (
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
)

const (
	start1 = 9
	start2 = 4
	//start1 = 4
	//start2 = 8
)

const (
	deterministicWin = 1000
	diracWin         = 21
)

func main() {
	s1, s2 := 0, 0
	p1, p2 := start1, start2
	d, count := 7, 0

	for {
		p1, s1, d, count = playDeterministic(p1, s1, d, count)
		if s1 >= deterministicWin {
			break
		}

		p2, s2, d, count = playDeterministic(p2, s2, d, count)
		if s2 >= deterministicWin {
			break
		}
	}
	fmt.Printf("part one: %v\n", count*aoc.Min(s1, s2))

	c1, c2 := playDirac(1, start1, start2, 0, 0)
	fmt.Printf("part two: %v\n", aoc.Max(c1, c2))
}

func play(p, s, d int) (int, int) {
	p += d
	if p > 10 {
		p -= 10
	}
	return p, s + p
}

func playDeterministic(p, s, d, c int) (int, int, int, int) {
	d--
	if d < 0 {
		d = 9
	}
	p, s = play(p, s, d)
	return p, s, d, c + 3
}

var dirac = diracDice()

func diracDice() [][2]int {
	hist := [10]int{}
	for t1 := 1; t1 <= 3; t1++ {
		for t2 := 1; t2 <= 3; t2++ {
			for t3 := 1; t3 <= 3; t3++ {
				hist[t1+t2+t3]++
			}
		}
	}

	var d [][2]int
	for i := 3; i <= 9; i++ {
		d = append(d, [2]int{i, hist[i]})
	}
	return d
}

func playDirac(player, p1, p2, s1, s2 int) (int, int) {
	sc1, sc2 := 0, 0
	var c1, c2 int

	for _, d := range dirac {
		if player == 1 {
			p, s := play(p1, s1, d[0])
			if s >= diracWin {
				c1, c2 = 1, 0
			} else {
				c1, c2 = playDirac(2, p, p2, s, s2)
			}
		} else {
			p, s := play(p2, s2, d[0])
			if s >= diracWin {
				c1, c2 = 0, 1
			} else {
				c1, c2 = playDirac(1, p1, p, s1, s)
			}
		}

		sc1 += d[1] * c1
		sc2 += d[1] * c2
	}
	return sc1, sc2
}
