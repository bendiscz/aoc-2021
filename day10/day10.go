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

var pairs = [255]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var errorScore = [255]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completeScore = [255]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	total := 0
	var scores []int
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		score1, score2 := score(line)
		total += score1
		if score2 > 0 {
			scores = append(scores, score2)
		}
	}

	sort.Ints(scores)

	fmt.Printf("part one: %v\n", total)
	fmt.Printf("part two: %v\n", scores[len(scores)/2])
}

func score(line string) (int, int) {
	var stack []byte
	for i := 0; i < len(line); i++ {
		ch := line[i]
		if pairs[ch] != 0 {
			stack = append(stack, pairs[ch])
		} else {
			if len(stack) == 0 {
				return errorScore[ch], 0
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if ch != top {
				return errorScore[ch], 0
			}
		}
	}

	complete := 0
	for i := len(stack) - 1; i >= 0; i-- {
		complete = complete*5 + completeScore[stack[i]]
	}

	return 0, complete
}
