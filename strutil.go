// Package stringutil contains utility functions for working with strings
package strutil


func Reverse(s1 string) string {
	l := len(s1)
	s2 := make([]rune, l)
	j:=0
	for i:=len(s1)-1; i>=0 && j< len(s1); i,j = i-1, j+1 {
		s2[i] = rune(s1[j])
	}
	return string(s2)
}
