package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"strings"
	"time"
)

//go:embed input1
var input []byte

type block struct {
	index    int
	sub, add int
}

type link struct {
	i1, i2 int
	diff   int
}

func (l link) max() (int, int) {
	if l.diff > 0 {
		return 9 - l.diff, 9
	} else {
		return 9, 9 + l.diff
	}
}

func (l link) min() (int, int) {
	if l.diff > 0 {
		return 1, 1 + l.diff
	} else {
		return 1 - l.diff, 1
	}
}

func digit(x int) byte {
	return byte('0' + x)
}

func readBlock(scanner *bufio.Scanner, index int) block {
	b := block{index: index}
	for i := 0; i < 18; i++ {
		scanner.Scan()
		switch i {
		case 5:
			b.sub = aoc.ParseInt(strings.Split(scanner.Text(), " ")[2])
		case 15:
			b.add = aoc.ParseInt(strings.Split(scanner.Text(), " ")[2])
		}
	}
	return b
}

func main() {
	t := time.Now()
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var stack []block
	var links []link
	for i := 0; i < 14; i++ {
		b := readBlock(scanner, i)
		if b.sub < 0 {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			links = append(links, link{p.index, b.index, p.add + b.sub})
		} else {
			stack = append(stack, b)
		}
	}

	var max, min [14]byte
	for _, l := range links {
		max1, max2 := l.max()
		max[l.i1], max[l.i2] = digit(max1), digit(max2)
		min1, min2 := l.min()
		min[l.i1], min[l.i2] = digit(min1), digit(min2)
	}

	fmt.Printf("part one: %s\n", string(max[:]))
	fmt.Printf("part two: %s\n", string(min[:]))
	fmt.Printf("time: %v\n", time.Since(t))
}
