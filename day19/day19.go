package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"math"
	"regexp"
	"strings"
)

//go:embed input1
var input []byte

var pattern = regexp.MustCompile(`^--- scanner (\d+) ---$`)

type xyz [3]int

var rotations = [...]xyz{
	// facing ±X
	{1, 2, 3},
	{1, -3, 2},
	{1, -2, -3},
	{1, 3, -2},
	{-1, -2, 3},
	{-1, 3, 2},
	{-1, 2, -3},
	{-1, -3, -2},
	// facing ±Y
	{2, 3, 1},
	{2, -1, 3},
	{2, -3, -1},
	{2, 1, -3},
	{-2, -3, 1},
	{-2, 1, 3},
	{-2, 3, -1},
	{-2, -1, -3},
	// facing ±Z
	{3, 1, 2},
	{3, -2, 1},
	{3, -1, -2},
	{3, 2, -1},
	{-3, -1, 2},
	{-3, 2, 1},
	{-3, 1, -2},
	{-3, -2, -1},
}

func (p xyz) rotate(r xyz) xyz {
	var p2 xyz
	for i, a := range r {
		if a < 0 {
			p[i] = -p[i]
		}
		p2[aoc.Abs(a)-1] = p[i]
	}
	return p2
}

func (p xyz) dist(p2 xyz) xyz {
	return xyz{
		p2[0] - p[0],
		p2[1] - p[1],
		p2[2] - p[2],
	}
}

func (p xyz) add(p2 xyz) xyz {
	return xyz{
		p2[0] + p[0],
		p2[1] + p[1],
		p2[2] + p[2],
	}
}

func (p xyz) manhattan(p2 xyz) int {
	return aoc.Abs(p2[0]-p[0]) + aoc.Abs(p2[1]-p[1]) + aoc.Abs(p2[2]-p[2])
}

type box struct {
	index  int
	points []xyz
	linked bool
	offset xyz
}

func (b *box) rotate(r xyz) {
	for i := 0; i < len(b.points); i++ {
		b.points[i] = b.points[i].rotate(r)
	}
}

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var boxes []*box
	for scanner.Scan() {
		m := pattern.FindStringSubmatch(scanner.Text())
		b := box{index: aoc.ParseInt(m[1])}
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				break
			}
			v := strings.Split(line, ",")
			b.points = append(b.points, xyz{
				aoc.ParseInt(v[0]),
				aoc.ParseInt(v[1]),
				aoc.ParseInt(v[2]),
			})
		}
		boxes = append(boxes, &b)
	}

	boxes[0].linked = true
	link(boxes, 0)

	beacons := map[xyz]struct{}{}
	for _, b := range boxes {
		for _, p := range b.points {
			beacons[p.add(b.offset)] = struct{}{}
		}
	}

	max := math.MinInt
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes); j++ {
			d := boxes[i].offset.manhattan(boxes[j].offset)
			if d > max {
				max = d
			}
		}
	}

	fmt.Printf("part one: %v\n", len(beacons))
	fmt.Printf("part two: %v\n", max)
}

func match(b0, b1 *box) (offset xyz, rotation xyz, ok bool) {
	for _, rotation = range rotations {
		m := map[xyz]int{}
		for _, p0 := range b0.points {
			for _, p1 := range b1.points {
				p1 = p1.rotate(rotation)
				d := p1.dist(p0)
				c := m[d] + 1
				if c == 12 {
					return d, rotation, true
				}
				m[d] = c
			}
		}
	}
	return
}

func link(boxes []*box, index int) {
	b0 := boxes[index]
	for i := 0; i < len(boxes); i++ {
		b := boxes[i]
		if b.linked {
			continue
		}
		if offset, rotation, ok := match(b0, b); ok {
			b.rotate(rotation)
			b.offset = b0.offset.add(offset)
			b.linked = true
			link(boxes, i)
		}
	}
}
