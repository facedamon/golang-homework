package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type NumChangeWithAtomic struct {
	num int32
	sync.WaitGroup
}

func main() {
	n := new(NumChangeWithAtomic)
	n.Add(10)

	fmt.Println("being...")

	for i := 0; i < 10; i++ {
		go func(index int) {
			defer n.Done()
			for j := 0; j < 1000; j++ {
				//n.num++
				atomic.AddInt32(&n.num, 1)
			}
		}(i)
	}

	n.Wait()
	fmt.Println("num=", n.num)
}
