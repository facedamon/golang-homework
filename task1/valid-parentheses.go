package main

import "fmt"

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	keys := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for _, v := range s {
		//如果匹配上key，则需要出栈
		if _, ok := keys[byte(v)]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != keys[byte(v)] {
				return false
			}
			//出栈
			stack = stack[:len(stack)-1]
		} else {
			//入栈
			stack = append(stack, byte(v))
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("(]"))     //false
	fmt.Println(isValid("([])"))   //true
	fmt.Println(isValid("([)]"))   //false
	fmt.Println(isValid("}{"))     //false
}
