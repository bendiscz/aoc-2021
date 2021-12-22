package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"regexp"
)

//go:embed input1
var input []byte

var pattern = regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)

const (
	S = 50
	D = 2*S + 1
)

type space [D][D][D]bool

func (s *space) count() (count int) {
	for x := 0; x < D; x++ {
		for y := 0; y < D; y++ {
			for z := 0; z < D; z++ {
				if s[x][y][z] {
					count++
				}
			}
		}
	}
	return
}

func main() {
	var c space
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		m := pattern.FindStringSubmatch(scanner.Text())
		x1, x2, cx := parseInterval(m[2], m[3])
		y1, y2, cy := parseInterval(m[4], m[5])
		z1, z2, cz := parseInterval(m[6], m[7])
		if cx || cy || cz {
			continue
		}
		b := m[1] == "on"

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				for z := z1; z <= z2; z++ {
					c[x][y][z] = b
				}
			}
		}
	}

	fmt.Printf("part one: %v\n", c.count())
}

func parseInterval(a, b string) (int, int, bool) {
	x, cx := clamp(aoc.ParseInt(a))
	y, cy := clamp(aoc.ParseInt(b))
	return x + S, y + S, cx || cy
}

func clamp(x int) (int, bool) {
	clamped := false
	if x < -S {
		x = -S
		clamped = true
	}
	if x > S {
		x = S
		clamped = true
	}
	return x, clamped
}
