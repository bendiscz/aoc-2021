package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"math"
	"regexp"
	"sort"
	"time"
)

//go:embed input1
var input []byte

var pattern = regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)

type box struct {
	a     [3]axis
	value bool
}

type axis struct {
	p, q int
}

type interval struct {
	p, q  int
	boxes map[int]struct{}
}

func (i *interval) size() int {
	return i.q - i.p + 1
}

func (i *interval) add(b int) {
	i.boxes[b] = struct{}{}
}

func (i *interval) has(b int) bool {
	_, ok := i.boxes[b]
	return ok
}

func (i *interval) copy() *interval {
	i2 := interval{
		i.p, i.q, make(map[int]struct{}),
	}
	for b := range i.boxes {
		i2.boxes[b] = struct{}{}
	}
	return &i2
}

func inf() *interval {
	return &interval{
		p:     math.MinInt,
		q:     math.MaxInt,
		boxes: map[int]struct{}{},
	}
}

func main() {
	var boxes []box
	var sxs = []*interval{inf()}
	var sys = []*interval{inf()}
	var szs = []*interval{inf()}

	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		m := pattern.FindStringSubmatch(scanner.Text())
		x1, x2 := aoc.ParseInt(m[2]), aoc.ParseInt(m[3])
		y1, y2 := aoc.ParseInt(m[4]), aoc.ParseInt(m[5])
		z1, z2 := aoc.ParseInt(m[6]), aoc.ParseInt(m[7])

		i := len(boxes)
		b := makeBox(x1, x2, y1, y2, z1, z2, m[1] == "on")
		boxes = append(boxes, b)

		sxs = insert(sxs, i, x1, x2)
		sys = insert(sys, i, y1, y2)
		szs = insert(szs, i, z1, z2)
	}

	t := time.Now()
	volume := 0
	for k, sx := range sxs {
		fmt.Printf("%d/%d\n", k, len(sxs))
		ax := map[int]struct{}{}
		for bx := range sx.boxes {
			ax[bx] = struct{}{}
		}
		if len(ax) == 0 {
			continue
		}

		for _, sy := range sys {
			var axy []int
			for by := range sy.boxes {
				if _, ok := ax[by]; ok {
					axy = append(axy, by)
				}
			}
			if len(axy) == 0 {
				continue
			}
			sort.Ints(axy)

			for _, sz := range szs {
				for i := len(axy) - 1; i >= 0; i-- {
					if sz.has(axy[i]) {
						if boxes[axy[i]].value {
							volume += sx.size() * sy.size() * sz.size()
						}
						break
					}
				}
			}
		}
	}

	fmt.Printf("part two: %v\n", volume)
	fmt.Printf("time: %v\n", time.Since(t))
}

func makeBox(x1, x2, y1, y2, z1, z2 int, value bool) box {
	return box{
		a: [...]axis{
			{x1, x2},
			{y1, y2},
			{z1, z2},
		},
		value: value,
	}
}

func insert(ss []*interval, b, p, q int) []*interval {
	var next []*interval
	for _, s := range ss {
		if p > s.p && p <= s.q {
			s2 := s.copy()
			s2.q = p - 1
			s.p = p
			next = append(next, s2)
		}

		if q >= s.p && q < s.q {
			s2 := s.copy()
			s2.q = q
			s.p = q + 1
			next = append(next, s2)
		}

		next = append(next, s)
	}

	for _, s := range next {
		if p <= s.p && q >= s.q {
			s.add(b)
		}
	}

	return next
}
