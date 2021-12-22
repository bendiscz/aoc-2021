package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"regexp"
	"time"
)

//go:embed input1
var input []byte

var pattern = regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)

type axis struct {
	p, q int
}

type box struct {
	a      [3]axis
	weight int
}

func (b box) volume() int { return b.a[0].size() * b.a[1].size() * b.a[2].size() }

func (a axis) size() int              { return a.q - a.p + 1 }
func (a axis) intersects(b axis) bool { return a.q >= b.p && a.p <= b.q }
func (a axis) small() bool            { return a.p >= -50 && a.q <= 50 }

func intersect(c1, c2 box) (box, bool) {
	for i := 0; i < 3; i++ {
		if !c1.a[i].intersects(c2.a[i]) {
			return box{}, false
		}
	}

	return box{
		a: [3]axis{
			{aoc.Max(c1.a[0].p, c2.a[0].p), aoc.Min(c1.a[0].q, c2.a[0].q)},
			{aoc.Max(c1.a[1].p, c2.a[1].p), aoc.Min(c1.a[1].q, c2.a[1].q)},
			{aoc.Max(c1.a[2].p, c2.a[2].p), aoc.Min(c1.a[2].q, c2.a[2].q)},
		},
		weight: 0,
	}, true
}

func main() {
	t := time.Now()
	partOneDone := false
	var boxes []box
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		m := pattern.FindStringSubmatch(scanner.Text())
		x1, x2 := aoc.ParseInt(m[2]), aoc.ParseInt(m[3])
		y1, y2 := aoc.ParseInt(m[4]), aoc.ParseInt(m[5])
		z1, z2 := aoc.ParseInt(m[6]), aoc.ParseInt(m[7])
		bc := box{a: [...]axis{{x1, x2}, {y1, y2}, {z1, z2}}, weight: 1}

		if !partOneDone && (!bc.a[0].small() || !bc.a[1].small() || !bc.a[2].small()) {
			fmt.Printf("part one: %v (time %v)\n", count(boxes), time.Since(t))
			partOneDone = true
		}

		for _, b := range boxes[:] {
			if bi, ok := intersect(bc, b); ok {
				bi.weight = -1 * b.weight
				boxes = append(boxes, bi)
			}
		}

		if m[1] == "on" {
			boxes = append(boxes, bc)
		}
	}
	fmt.Printf("part two: %v (time %v)\n", count(boxes), time.Since(t))
}

func count(boxes []box) (volume int) {
	for _, b := range boxes {
		volume += b.weight * b.volume()
	}
	return
}
