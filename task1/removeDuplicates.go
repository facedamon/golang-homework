package main

import "fmt"

// 采用快慢指针法
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 3, 3}))
}
