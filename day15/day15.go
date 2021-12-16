package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
)

//go:embed input1
var input []byte

var board [500][500]int
var d int

var dist [500][500]int
var visited [500][500]bool

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		d = len(line)
		for c := 0; c < d; c++ {
			board[row][c] = int(line[c] - '0')
		}
		row++
	}

	reset()
	dijkstra()

	fmt.Printf("part one: %v\n", dist[d-1][d-1])

	expand()
	reset()
	dijkstra()
	fmt.Printf("part two: %v\n", dist[d-1][d-1])
}

func findMin() (int, int) {
	mr, mc, p := -1, -1, math.MaxInt
	for r := 0; r < d; r++ {
		for c := 0; c < d; c++ {
			if !visited[r][c] && dist[r][c] < p {
				p = dist[r][c]
				mr, mc = r, c
			}
		}
	}

	if mr != -1 {
		visited[mr][mc] = true
	}
	return mr, mc
}

func dijkstra() {
	for {
		mr, mc := findMin()
		if mr < 0 || mr == d-1 && mc == d-1 {
			break
		}

		if mr < d-1 {
			r, c := mr+1, mc
			a := dist[mr][mc] + board[r][c]
			if a < dist[r][c] {
				dist[r][c] = a
			}
		}
		if mr > 0 {
			r, c := mr-1, mc
			a := dist[mr][mc] + board[r][c]
			if a < dist[r][c] {
				dist[r][c] = a
			}
		}
		if mc < d-1 {
			r, c := mr, mc+1
			a := dist[mr][mc] + board[r][c]
			if a < dist[r][c] {
				dist[r][c] = a
			}
		}
		if mc > 0 {
			r, c := mr, mc-1
			a := dist[mr][mc] + board[r][c]
			if a < dist[r][c] {
				dist[r][c] = a
			}
		}
	}
}

func reset() {
	for r := 0; r < d; r++ {
		for c := 0; c < d; c++ {
			dist[r][c] = 10000000000
			visited[r][c] = false
		}
	}
	dist[0][0] = 0
}

func expand() {
	for r := 0; r < d; r++ {
		for c := 0; c < d; c++ {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					v := (board[r][c]-1+i+j)%9 + 1
					board[i*d+r][j*d+c] = v
				}
			}
		}
	}
	d *= 5
}
