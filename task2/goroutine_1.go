package main

import (
	"fmt"
	"sync"
)

var wg = new(sync.WaitGroup)

func printNum(flag bool) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if flag {
			if i%2 == 0 {
				fmt.Println("偶数协程A---:", i)
			}
		} else {
			if i%2 != 0 {
				fmt.Println("奇数协程B---:", i)
			}
		}
	}
}

func main() {

	wg.Add(2)

	go printNum(true)
	go printNum(false)

	wg.Wait()
}
