// https://leetcode.com/problems/valid-number
package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	StartState = iota
	NumState
	DotState
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}

	if len(s) > 1 {
		switch s[len(s)-1] {
		case '+', '-':
			return false
		}
	}

	i, state := 0, StartState
	nDigits, eCount := 0, 0
	for i < len(s) {
		switch state {
		case StartState:
			switch s[i] {
			case '-', '+':
				state = NumState
				i++
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				state = NumState
			case '.':
				if eCount > 0 {
					return false
				}
				state = DotState
				i++
			default:
				return false
			}
		case NumState:
			switch s[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				nDigits++
				i++
			case 'e':
				if eCount > 0 || nDigits == 0 {
					return false
				}
				state = StartState
				eCount++
				i++
			case '.':
				if eCount > 0 {
					return false
				}
				state = DotState
				i++
			default:
				return false
			}
		case DotState:
			switch s[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				nDigits++
				i++
			case 'e':
				if eCount > 0 || nDigits == 0 {
					return false
				}
				eCount++
				state = StartState
				i++
			default:
				return false
			}
		}
	}
	return (state == NumState) || (state == DotState && nDigits > 0)
}

func main() {
	cases := []struct {
		expr string
		val  bool
	}{
		{"4e+", false},
		{"1e.", false},
		{".e1", false},
		{" -90e3   ", true},
		{" 99e2.5 ", false},
		{"0", true},
		{" 0.1 ", true},
		{"abc", false},
		{"1 a", false},
		{"2e10", true},
		{" 1e", false},
		{"e3", false},
		{" 6e-1", true},
		{"53.5e93", true},
		{" --6 ", false},
		{"-+3", false},
		{"95a54e53", false},
		{".1", true},
		{"3.", true},
		{".", false},
	}

	ok := true
	for _, tc := range cases {
		val := isNumber(tc.expr)
		err := ""
		if val != tc.val {
			err = " ERROR"
			ok = false
		}
		fmt.Printf("%q, expected=%v, got=%v%s\n", tc.expr, tc.val, val, err)
		//break
	}

	if !ok {
		os.Exit(1)
	}
}
