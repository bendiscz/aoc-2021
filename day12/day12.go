package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input1
var input []byte

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		p := strings.IndexByte(line, '-')
		n1, n2 := line[:p], line[p+1:]
		addEdge(n1, n2)
		addEdge(n2, n1)
	}

	fmt.Printf("part one: %v\n", travel(newPath(1), "start"))
	fmt.Printf("part two: %v\n", travel(newPath(2), "start"))
}

var edges = map[string][]string{}

type path struct {
	part    int
	visited map[string]bool
	second  string
}

func newPath(part int) *path {
	return &path{
		part:    part,
		visited: map[string]bool{},
	}
}

func (p *path) enter(n string) bool {
	if !small(n) {
		return true
	}

	if p.visited[n] {
		if p.part == 1 {
			return false
		}
		if n == "start" || n == "end" || len(p.second) > 0 {
			return false
		}
		p.second = n
	} else {
		p.visited[n] = true
	}
	return true
}

func (p *path) leave(n string) {
	if n == p.second {
		p.second = ""
	} else {
		delete(p.visited, n)
	}
}

func small(s string) bool {
	for _, r := range s {
		return unicode.IsLower(r)
	}
	return false
}

func addEdge(m, n string) {
	e := edges[m]
	e = append(e, n)
	edges[m] = e
}

func travel(p *path, n string) int {
	if n == "end" {
		return 1
	}

	if !p.enter(n) {
		return 0
	}
	defer p.leave(n)

	sum := 0
	for _, m := range edges[n] {
		sum += travel(p, m)
	}
	return sum
}
