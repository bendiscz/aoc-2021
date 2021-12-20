package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input1
var input []byte

const (
	D = 100 + 2*O
	O = 50
)

type image [D][D]bool

func (img *image) get(x0, y0 int) int {
	v := 0
	for y := y0 - 1; y <= y0+1; y++ {
		for x := x0 - 1; x <= x0+1; x++ {
			var p bool
			if x < 0 || x >= D || y < 0 || y >= D {
				p = img[0][0]
			} else {
				p = img[x][y]
			}

			v *= 2
			if p {
				v++
			}
		}
	}
	return v
}

func (img *image) count() int {
	count := 0
	for x := 0; x < D; x++ {
		for y := 0; y < D; y++ {
			if img[x][y] {
				count++
			}
		}
	}
	return count
}

func (img *image) enhance(key string) *image {
	var img2 image
	for x := 0; x < D; x++ {
		for y := 0; y < D; y++ {
			img2[x][y] = key[img.get(x, y)] == '#'
		}
	}
	return &img2
}

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Scan()
	key := scanner.Text()
	scanner.Scan()

	img := &image{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			img[x+O][y+O] = line[x] == '#'
		}
		y++
	}

	img = img.enhance(key)
	img = img.enhance(key)
	fmt.Printf("part one: %v\n", img.count())

	for i := 0; i < 48; i++ {
		img = img.enhance(key)
	}
	fmt.Printf("part two: %v\n", img.count())
}
