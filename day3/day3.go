package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input1
var input []byte

func main() {
	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("part one: %v\n", solvePartOne(lines))
	fmt.Printf("part two: %v\n", solvePartTwo(lines))
}

func solvePartOne(lines []string) int {
	var bits [63]int
	lastBit := 0

	for _, num := range lines {
		i := 0
		for i < len(num) {
			if num[i] == '1' {
				bits[i]++
			} else {
				bits[i]--
			}
			i++
		}
		if i > lastBit {
			lastBit = i
		}
	}

	g, e := 0, 0
	for i := 0; i < lastBit; i++ {
		g, e = 2*g, 2*e
		if bits[i] >= 0 {
			g++
		} else {
			e++
		}
	}
	return g * e
}

func solvePartTwo(lines []string) int {
	oxy, _ := strconv.ParseInt(find(lines, false), 2, 64)
	co2, _ := strconv.ParseInt(find(lines, true), 2, 64)
	return int(oxy) * int(co2)
}

func find(lines []string, inv bool) string {
	filter := make([]bool, len(lines))

	for pos := 0; pos < len(lines[0]); pos++ {
		bits := 0
		for i, num := range lines {
			if filter[i] {
				continue
			}
			if num[pos] == '1' {
				bits++
			} else {
				bits--
			}
		}

		var bit byte
		cond := bits >= 0
		if inv {
			cond = !cond
		}
		if cond {
			bit = '1'
		} else {
			bit = '0'
		}

		selected, count := -1, 0
		for i, num := range lines {
			if filter[i] {
				continue
			}

			if num[pos] == bit {
				selected = i
				count++
			} else {
				filter[i] = true
			}
		}

		if count == 1 {
			return lines[selected]
		}
	}

	return ""
}
