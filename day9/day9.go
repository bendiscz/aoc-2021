package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"sort"
)

//go:embed input1
var input []byte

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = int(line[i] - '0')
		}
		board = append(board, row)
	}
	R, C = len(board), len(board[0])

	var lows []RC

	risk := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			l := board[r][c]
			if r > 0 && board[r-1][c] <= l {
				continue
			}
			if r < R-1 && board[r+1][c] <= l {
				continue
			}
			if c > 0 && board[r][c-1] <= l {
				continue
			}
			if c < C-1 && board[r][c+1] <= l {
				continue
			}
			risk += l + 1
			lows = append(lows, RC{r, c})
		}
	}

	fmt.Printf("part one: %v\n", risk)

	mask = make([][]bool, R)
	for r := range mask {
		mask[r] = make([]bool, C)
	}

	var basins []int
	for _, low := range lows {
		basins = append(basins, fill(low))
	}

	sort.Ints(basins)
	if len(basins) >= 3 {
		basins = basins[len(basins)-3:]
		fmt.Printf("part two: %v\n", basins[0]*basins[1]*basins[2])
	}
}

type RC struct {
	r, c int
}

var R, C int
var board [][]int
var mask [][]bool

func fill(low RC) int {
	count := 0
	stack := []RC{low}
	for len(stack) > 0 {
		rc := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if mask[rc.r][rc.c] {
			continue
		}
		if board[rc.r][rc.c] == 9 {
			continue
		}

		count++
		mask[rc.r][rc.c] = true

		if rc.r > 0 {
			stack = append(stack, RC{rc.r - 1, rc.c})
		}
		if rc.r < R-1 {
			stack = append(stack, RC{rc.r + 1, rc.c})
		}
		if rc.c > 0 {
			stack = append(stack, RC{rc.r, rc.c - 1})
		}
		if rc.c < C-1 {
			stack = append(stack, RC{rc.r, rc.c + 1})
		}
	}
	return count
}
