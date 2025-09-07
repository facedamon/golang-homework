package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	//convert to string
	s := strconv.Itoa(x)
	s2 := make([]string, len(s))
	// 倒序
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("%c", s[i])
		s2 = append(s2, string(s[i]))
	}

	fmt.Printf("\ns=%s, s2=%s\n", s, strings.Join(s2, ""))

	return s == strings.Join(s2, "")
}

func main() {
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(1234))
}
