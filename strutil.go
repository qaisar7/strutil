// Package stringutil contains utility functions for working with strings
package strutil

import (
	"fmt"
	"strings"
)

const (
	space rune = rune(' ')
)

// Reverse reverses a given string.
// Hello -> olleH
func Reverse(s1 string) (string, error) {
	if s1 == "" {
		return "", fmt.Errorf("strutil/Reverse(): the provided string is empty")
	}
	l := len(s1)
	s2 := make([]rune, l)
	j := 0
	for i := len(s1) - 1; i >= 0 && j < len(s1); i, j = i-1, j+1 {
		s2[i] = rune(s1[j])
	}
	return string(s2), nil
}

// CamelCase changes a given string to camelCase by removing spaces and capitalizing the next rune after space.
// 'hey mate' -> heyMate
func CamelCase(s string) string {
	s2 := []rune{}
	c := false
	for _, r := range s {
		if r == space {
			c = true
			continue
		}
		if c {
			r = rune(strings.ToUpper(string(r))[0])
			c = false
		}
		s2 = append(s2, r)
	}
	return string(s2)
}

func main() {
	fmt.Println(CamelCase("hey mate"))
}
