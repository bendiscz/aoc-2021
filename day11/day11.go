package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input1
var input []byte

var board [10][10]int

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < 10; x++ {
			board[x][y] = int(line[x] - '0')
		}
		y++
	}

	count := 0
	all := -1
	for i := 0; ; i++ {
		glow()
		for iterate() {
		}
		x := countFlashes()
		if i < 100 {
			count += x
		}
		if x == 100 {
			all = i + 1
			break
		}
	}

	fmt.Printf("part one: %v\n", count)
	fmt.Printf("part two: %v\n", all)
}

func glow() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			board[x][y]++
		}
	}
}

func iterate() bool {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if board[x][y] > 9 {
				flash(x, y)
				return true
			}
		}
	}
	return false
}

func flash(x0, y0 int) {
	board[x0][y0] = -1
	for x := x0 - 1; x <= x0+1; x++ {
		for y := y0 - 1; y <= y0+1; y++ {
			if x < 0 || x >= 10 || y < 0 || y >= 10 {
				continue
			}
			if board[x][y] >= 0 {
				board[x][y]++
			}
		}
	}
}

func countFlashes() int {
	count := 0
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if board[x][y] == -1 {
				count++
				board[x][y] = 0
			}
		}
	}
	return count
}
