package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"strings"
)

type board struct {
	num  [5][5]int
	hit  [5][5]bool
	done bool
}

func (b *board) mark(num int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.num[i][j] == num {
				b.hit[i][j] = true
			}
		}
	}
}

func (b *board) check() bool {
	for i := 0; i < 5; i++ {
		r, c := true, true
		for j := 0; j < 5; j++ {
			if !b.hit[i][j] {
				r = false
			}
			if !b.hit[j][i] {
				c = false
			}
		}
		if r || c {
			return true
		}
	}

	return false
}

func (b *board) score() int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.hit[i][j] {
				score += b.num[i][j]
			}
		}
	}
	return score
}

//go:embed input1
var input []byte

func main() {
	numbers, boards := load()
	var firstScore, lastScore int

	for _, num := range numbers {
		for _, b := range boards {
			if b.done {
				continue
			}

			b.mark(num)
			if b.check() {
				b.done = true
				score := b.score() * num
				if firstScore == 0 {
					firstScore = score
				}
				lastScore = score
			}
		}
	}

	fmt.Printf("part one: %v\n", firstScore)
	fmt.Printf("part two: %v\n", lastScore)
}

func load() (numbers []int, boards []*board) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	if !scanner.Scan() {
		return
	}

	for _, s := range strings.Split(scanner.Text(), ",") {
		numbers = append(numbers, aoc.ParseInt(s))
	}

	for scanner.Scan() {
		var b board
		for i := 0; i < 5; i++ {
			if !scanner.Scan() {
				return
			}

			for j, s := range strings.Fields(scanner.Text()) {
				b.num[i][j] = aoc.ParseInt(s)
			}
		}

		boards = append(boards, &b)
	}
	return
}
