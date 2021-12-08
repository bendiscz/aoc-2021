package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

//go:embed input1
var input []byte

var digits = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

var perms [][7]int

func init() {
	var perm [7]int
	genPerms(perm, 0)
	fmt.Printf("perms: %d\n", len(perms))
}

func main() {

	//input = []byte("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf\n")

	scanner := bufio.NewScanner(bytes.NewReader(input))
	count1, count2 := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return unicode.IsSpace(r) || r == '|'
		})

		for _, f := range fields[10:] {
			if len(f) == 2 || len(f) == 3 || len(f) == 4 || len(f) == 7 {
				count1++
			}
		}

		if x, ok := findMatch(fields); ok {
			count2 += x
		}
	}

	fmt.Printf("part one: %v\n", count1)
	fmt.Printf("part two: %v\n", count2)
}

func mapDigit(digit string, perm [7]int) (int, bool) {
	var mapped []byte
	for i := 0; i < len(digit); i++ {
		d := perm[digit[i]-'a'] + 'a'
		mapped = append(mapped, byte(d))
	}

	sort.Slice(mapped, func(i, j int) bool {
		return mapped[i] < mapped[j]
	})

	x, ok := digits[string(mapped)]
	return x, ok
}

func findMatch(fields []string) (int, bool) {
loop:
	for _, perm := range perms {
		count := 0
		for i, f := range fields {
			x, ok := mapDigit(f, perm)
			if !ok {
				continue loop
			}
			if i >= 10 {
				count = count*10 + x
			}
		}
		return count, true
	}
	return 0, false
}

func genPerms(perm [7]int, index int) {
	if index == 7 {
		var hit [7]int
		copy(hit[:], perm[:])
		perms = append(perms, hit)
		return
	}

loop:
	for d := 0; d < 7; d++ {
		for i := 0; i < index; i++ {
			if perm[i] == d {
				continue loop
			}
		}

		perm[index] = d
		genPerms(perm, index+1)
	}
}
