// https://leetcode.com/problems/basic-calculator
package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	tokRe = regexp.MustCompile(`([()+-])`)
)

func tokenize(code string) []string {
	code = tokRe.ReplaceAllString(code, " $1 ")
	return strings.Fields(code)
}

func eval_one(tokens []string) (int, int, error) {
	// fmt.Printf("eval_one: %v\n", tokens)
	if len(tokens) == 0 {
		return 0, 0, fmt.Errorf("no tokens")
	}

	if tokens[0] == "(" { // (expr)
		count := 1
		i := 1
		for _, tok := range tokens[1:] {
			switch tok {
			case ")":
				count--
			case "(":
				count++
			}
			if count == 0 {
				break
			}
			i++
		}
		if count != 0 {
			return 0, 0, fmt.Errorf("unbalanced parenthesis")
		}
		val, err := eval(tokens[1:i])
		if err != nil {
			return 0, 0, err
		}
		return val, i + 1, nil
	}

	i, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, 0, err
	}
	return int(i), 1, nil
}

func calc(lval, rval int, op string) int {
	if op == "+" {
		return lval + rval
	}
	return lval - rval
}

func eval(tokens []string) (int, error) {
	lval, op, rval := 0, "+", 0
	for {
		// fmt.Printf("eval: %v\n", tokens)
		val, i, err := eval_one(tokens)
		if err != nil {
			return 0, err
		}
		rval = val
		if i >= len(tokens) {
			break
		}
		lval = calc(lval, rval, op)
		op = tokens[i]
		if op != "+" && op != "-" {
			return 0, fmt.Errorf("unknown op - %q", op)
		}
		tokens = tokens[i+1:]
	}

	return calc(lval, rval, op), nil
}

func calculate(s string) int {
	tokens := tokenize(s)
	val, _ := eval(tokens)
	return val
}

// expr -> ( expr )
// expr -> expr OP expr
// exprt -> NUM
// OP -> + | -

func main() {
	var cases = []struct {
		code  string
		value int
	}{
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
	}
	for _, tc := range cases {
		tokens := tokenize(tc.code)
		val, err := eval(tokens)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("code=%q, expected=%d, got=%d\n", tc.code, tc.value, val)
	}
}
