package main

import (
	"fmt"
	"sync"
)

type NumChange struct {
	sync.Mutex
	sync.WaitGroup
	num int
}

func main() {
	n := new(NumChange)
	n.Add(10)

	fmt.Println("begin...")

	for i := 0; i < 10; i++ {
		go func(index int) {
			defer n.Done()
			defer n.Unlock()
			n.Lock()
			for j := 0; j < 1000; j++ {
				n.num++
			}
		}(i)
	}
	n.Wait()
	fmt.Println("num=", n.num)
}
