package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"time"
)

//go:embed input1
var input []byte

type sea struct {
	rows [][]byte
	w, h int
}

func main() {
	s := &sea{}
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		s.rows = append(s.rows, []byte(scanner.Text()))
	}
	s.h, s.w = len(s.rows), len(s.rows[0])

	t := time.Now()
	count := 1
	for s.move() {
		count++
	}

	fmt.Printf("part one: %v (time %v)\n", count, time.Since(t))
	fmt.Printf("part two: %v\n", "Merry Christmas!")
}

func (s *sea) move() (moved bool) {
	moved = s.moveEast() || moved
	moved = s.moveSouth() || moved
	return
}

func (s *sea) moveEast() (moved bool) {
	for r := 0; r < s.h; r++ {
		row := s.rows[r]
		c, free := 1, row[0] == '.'
		for c < s.w {
			if row[c] == '.' && row[c-1] == '>' {
				moved = true
				row[c-1], row[c] = '.', '>'
				c++
			}
			c++
		}
		if c == s.w && free && row[c-1] == '>' {
			moved = true
			row[c-1], row[0] = '.', '>'
		}
	}
	return
}

func (s *sea) moveSouth() (moved bool) {
	for c := 0; c < s.w; c++ {
		r, free := 1, s.rows[0][c] == '.'
		for r < s.h {
			if s.rows[r][c] == '.' && s.rows[r-1][c] == 'v' {
				moved = true
				s.rows[r-1][c], s.rows[r][c] = '.', 'v'
				r++
			}
			r++
		}
		if r == s.h && free && s.rows[r-1][c] == 'v' {
			moved = true
			s.rows[r-1][c], s.rows[0][c] = '.', 'v'
		}
	}
	return
}
