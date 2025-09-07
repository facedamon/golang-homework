package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	total := len(strs)
	/*fmt.Printf("%T\n", first)
	for _, i2 := range first {
		fmt.Printf("%T\n", i2) //int32 如果使用普通for循环则是uint8
		fmt.Printf("%s", string(i2))
	}*/
	for i, v := range first {
		//第一与第二开始对比
		for j := 1; j < total; j++ {
			//如果到尾部了 或者 第j行第i列的元素和第一行第i列元素不一样，则退出
			if i == len(strs[j]) || string(v) != string(strs[j][i]) {
				//退出时，最长前缀字符应是第i列之前的元素
				return first[:i]
			}
		}
	}
	return first
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
