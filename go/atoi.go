// https://leetcode.com/problems/string-to-integer-atoi/
package main

import (
	"fmt"
	"math"
)

func myAtoi(text string) int {
	if len(text) == 0 {
		return 0
	}
	// strip leading white space
	start := 0
	for i, c := range text {
		if c == ' ' || c == '\t' || c == '\n' {
			continue
		}
		start = i
		break
	}

	// ± leader
	mul := 1
	switch text[start] {
	case '-':
		mul = -1
		start++
	case '+':
		start++
	}

	n := 0
	for _, c := range text[start:] {
		if c < '0' || c > '9' {
			break
		}
		i := int(c - '0')
		n = n*10 + i

		if mul*n < math.MinInt32 {
			return math.MinInt32
		}
		if mul*n > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return mul * n
}

func main() {
	cases := []string{
		"42",
		"   -42",
		"4193 with words",
		"words and 987",
		"-91283472332",
	}
	for _, text := range cases {
		n := myAtoi(text)
		fmt.Printf("%q → %d\n", text, n)
	}
}
