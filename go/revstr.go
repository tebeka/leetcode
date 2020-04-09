// https://leetcode.com/problems/reverse-words-in-a-string/
package main

import (
	"fmt"
)

func split(s string) []string {
	var words []string

	start := 0
	for start < len(s) {
		// white space at beginning
		for ; start < len(s) && s[start] == ' '; start++ {
		}
		if start >= len(s) {
			break
		}

		// run until whitespace/end
		end := start + 1
		for ; end < len(s); end++ {
			if s[end] == ' ' {
				break
			}
		}
		words = append(words, s[start:end])
		start = end + 1
	}

	return words
}

func numChars(words []string) int {
	size := 0
	for _, w := range words {
		size += len(w)
	}
	return size
}

func revJoin(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// num chars + num words - 1 spaces
	out := make([]byte, numChars(words)+len(words)-1)
	n := 0
	for i := len(words) - 1; i > 0; i-- {
		copy(out[n:], words[i][:])
		out[n+len(words[i])] = ' '
		n += len(words[i]) + 1
	}
	// add first word last
	copy(out[n:], words[0][:])
	return string(out)
}

func reverseWords(s string) string {
	words := split(s)
	return revJoin(words)
}

func main() {
	testCases := []struct {
		in  string
		out string
	}{
		{"  hello world!  ", "world! hello"},
		{"the sky is blue", "blue is sky the"},
		{"a good   example", "example good a"},
	}

	for _, tc := range testCases {
		out := reverseWords(tc.in)
		if tc.out == out {
			continue
		}
		fmt.Printf("%q: expected %q, got %q\n", tc.in, tc.out, out)
	}
}
