package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/bendiscz/aoc-2021"
	"regexp"
	"strings"
)

//go:embed input1
var input []byte

var pattern = regexp.MustCompile(`^fold along ([xy])=(\d+)$`)

const D = 2000

var paper = [D][D]bool{}

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		p := strings.IndexByte(line, ',')
		x, y := aoc.ParseInt(line[:p]), aoc.ParseInt(line[p+1:])
		paper[x][y] = true
	}

	//printPaper(20, 20)

	first := true
	for scanner.Scan() {
		m := pattern.FindStringSubmatch(scanner.Text())
		switch m[1] {
		case "x":
			foldX(aoc.ParseInt(m[2]))
		case "y":
			foldY(aoc.ParseInt(m[2]))
		}
		if first {
			fmt.Printf("part one: %v\n", count())
			first = false
		}
	}

	printPaper(40, 6)
}

func foldX(x0 int) {
	for y := 0; y < D; y++ {
		for x := x0; x < D; x++ {
			mx := 2*x0 - x
			if mx >= 0 {
				paper[mx][y] = paper[mx][y] || paper[x][y]
			}
			paper[x][y] = false
		}
	}
}

func foldY(y0 int) {
	for x := 0; x < D; x++ {
		for y := y0; y < D; y++ {
			my := 2*y0 - y
			//fmt.Printf("mapping %d,%d (%b) to %d,%d\n", x, y, paper[x][y], x, my)
			if my >= 0 {
				paper[x][my] = paper[x][my] || paper[x][y]
			}
			paper[x][y] = false
		}
		//printPaper(20, 20)
		//fmt.Printf("---\n")
	}
}

func count() int {
	c := 0
	for x := 0; x < D; x++ {
		for y := 0; y < D; y++ {
			if paper[x][y] {
				c++
			}
		}
	}
	return c
}

func printPaper(w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ch := '.'
			if paper[x][y] {
				ch = '#'
			}
			fmt.Printf("%c", ch)
		}
		fmt.Printf("\n")
	}
}
