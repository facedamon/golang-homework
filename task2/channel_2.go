package main

import (
	"fmt"
	"sync"
)

var wg4 = new(sync.WaitGroup)

func outputNumWithCache(c chan<- int) {
	defer close(c)
	for i := 1; i <= 100; i++ {
		c <- i
	}
}

func getNumWithCache(c <-chan int) {
	defer wg4.Done()
	for i := range c {
		fmt.Println("读取有缓冲通道数值=", i)
	}
}

func main() {
	c := make(chan int, 10)
	fmt.Println("begin...")
	wg4.Add(1)
	go outputNumWithCache(c)
	go getNumWithCache(c)
	wg4.Wait()
}
