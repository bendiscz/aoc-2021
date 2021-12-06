package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"strconv"
)

//go:embed input1
var input []byte

func main() {
	fmt.Printf("part one: %v\n", partOne(bufio.NewScanner(bytes.NewReader(input))))
	fmt.Printf("part two: %v\n", partTwo(bufio.NewScanner(bytes.NewReader(input))))
}

func partOne(scanner *bufio.Scanner) int {
	prev := math.MaxInt
	result := 0
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		if depth > prev {
			result++
		}

		prev = depth
	}

	return result
}

func partTwo(scanner *bufio.Scanner) int {
	win := [3]int{}
	i := 0
	for i < 3 && scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		win[i] = depth
		i++
	}

	if i < 3 {
		return 0
	}

	winA, winB, result := 0, 0, 0
	for scanner.Scan() {
		depth, _ := strconv.Atoi(scanner.Text())
		winA = win[0] + win[1] + win[2]
		winB = win[1] + win[2] + depth

		if winB > winA {
			result++
		}

		win[0], win[1], win[2] = win[1], win[2], depth
	}

	return result
}
