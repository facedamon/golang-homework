package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//key:值 value:下标
	m := make(map[int]int)
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 5}, 9))
}
