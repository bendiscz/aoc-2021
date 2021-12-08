package aoc

import (
	"crypto"
	_ "crypto/md5"
	"strconv"
)

type XY struct {
	x, y int
}

func ParseInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func MD5(data []byte) []byte {
	md5 := crypto.MD5.New()
	md5.Write(data)
	return md5.Sum(nil)
}
