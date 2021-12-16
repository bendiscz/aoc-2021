package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"math"
)

//go:embed input1
var input []byte

type bits struct {
	b []int
	p int
}

func (b *bits) read(n int) int {
	x := 0
	for n > 0 && b.p < len(b.b) {
		x = x*2 + b.b[b.p]
		b.p++
		n--
	}
	return x
}

func (b *bits) pos() int {
	return b.p
}

func main() {
	s := bytes.TrimSpace(input)
	var b bits
	for _, ch := range s {
		var x int
		if ch >= 'A' {
			x = int(ch - 'A' + 10)
		} else {
			x = int(ch - '0')
		}

		b.b = append(b.b, (x&8)>>3, (x&4)>>2, (x&2)>>1, x&1)
	}

	v, r := readPacket(&b)
	fmt.Printf("part one: %v\n", v)
	fmt.Printf("part two: %v\n", r)
}

func readPacket(b *bits) (vsum int, result int) {
	v := b.read(3)
	t := b.read(3)

	if t == 4 {
		return v, readLiteral(b)
	}

	vsum, results := readContainer(b)

	switch t {
	case 0:
		for _, x := range results {
			result += x
		}

	case 1:
		result = 1
		for _, x := range results {
			result *= x
		}

	case 2:
		result = math.MaxInt
		for _, x := range results {
			if x < result {
				result = x
			}
		}

	case 3:
		result = math.MinInt
		for _, x := range results {
			if x > result {
				result = x
			}
		}

	case 5:
		if len(results) == 2 && results[0] > results[1] {
			result = 1
		}

	case 6:
		if len(results) == 2 && results[0] < results[1] {
			result = 1
		}

	case 7:
		if len(results) == 2 && results[0] == results[1] {
			result = 1
		}
	}

	return vsum + v, result
}

func readLiteral(b *bits) (result int) {
	for {
		g := b.read(5)
		result = result*16 + (g & 15)
		if g&16 == 0 {
			break
		}
	}
	return
}

func readContainer(b *bits) (vsum int, results []int) {
	l := b.read(1)
	if l == 0 {
		return readContainer0(b)
	} else {
		return readContainer1(b)
	}
}

func readContainer0(b *bits) (vsum int, results []int) {
	bl := b.read(15)
	p := b.p
	for b.p-p < bl {
		v, r := readPacket(b)
		vsum += v
		results = append(results, r)
	}
	return
}

func readContainer1(b *bits) (vsum int, results []int) {
	pl := b.read(11)
	for i := 0; i < pl; i++ {
		v, r := readPacket(b)
		vsum += v
		results = append(results, r)
	}
	return vsum, results
}
