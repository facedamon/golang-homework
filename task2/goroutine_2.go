package main

import (
	"fmt"
	"sync"
	"time"
)

type jobs []int

var wg2 = new(sync.WaitGroup)

func executeJob(b jobs) {
	if len(b) == 0 {
		return
	}
	wg2.Add(len(b))
	for i, job := range b {
		go func(index int) {
			defer wg2.Done()
			start := time.Now()
			//模拟执行任务
			time.Sleep(2 * time.Second)
			fmt.Println("任务", i, "任务内容=", job, "执行时间=", time.Now().Sub(start).Seconds())
		}(i)
	}
}

func main() {
	bs := jobs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	executeJob(bs)
	wg2.Wait()
}
