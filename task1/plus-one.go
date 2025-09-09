package main

import (
	"fmt"
)

//第一个版本思路不对
//func plusOne(digits []int) []int {
//	var s string
//	for _, digit := range digits {
//		s += strconv.Itoa(digit)
//	}
//	//这里不对，会超长 没有big decimal的库
//	b, err := strconv.ParseInt(s, 10, 64)
//	if err != nil {
//		_ = fmt.Errorf("%w", err)
//	}
//	b++
//	ss := strconv.FormatInt(b, 10)
//	r := make([]int, 0)
//	for i := 0; i < len(ss); i++ {
//		st, _ := strconv.Atoi(string(ss[i]))
//		r = append(r, st)
//	}
//	return r
//}

// 只需要判断有没有进位
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		//如果进位则digits[i]=0,否则没有进位直接结束
		digits[i] = digits[i] % 10
		if digits[i]%10 != 0 {
			return digits
		}
	}
	//999这种情况
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9, 9, 9, 9, 9, 9, 9}))
	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 9}))
}
