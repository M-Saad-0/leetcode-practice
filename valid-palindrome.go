package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	pattern := regexp.MustCompile("[a-zA-Z0-9]")
	strs := pattern.FindAllString(s, -1)
	s = strings.Join(strs, "")
	s = strings.ToLower(s)
	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(s)/2; i++ {
		if s[len(s)-i-1] != s[i] {
			return false
		}
	}
	return true
}
