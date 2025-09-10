package main

import "fmt"

func changeValue(num *int) {
	*num += 10
}

func changeSlice(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}

func main() {
	num := 1
	changeValue(&num)
	fmt.Println(num)

	nums := []int{1, 2}
	changeSlice(nums)
	fmt.Println(nums)
}
