package main

import "fmt"

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	for k, vv := range m {
		if vv == 1 {
			return k
		}
	}
	return 0
}

func main() {
	i := singleNumber([]int{4, 1, 2, 1, 2})
	fmt.Printf("i=%d\n", i)
}
