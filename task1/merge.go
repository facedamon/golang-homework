package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return make([][]int, 0)
	}
	sort.Slice(intervals, func(i, j int) bool {
		//按照每行第一个元素排序
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	for _, interval := range intervals {
		left := interval[0]
		right := interval[1]
		//如果当前区间的left比merge中最后一个区间的right大，表示没有重合
		if len(merged) == 0 || merged[len(merged)-1][1] < left {
			merged = append(merged, []int{left, right})
		} else {
			//重合。需要用当前区间的right更新merge中最后一个区间的right，两个right取max
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], right)
		}
	}
	return merged
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{4, 7}, {1, 5}}))
}
