package main

import (
	"fmt"
	"sync"
)

var wg3 = new(sync.WaitGroup)

func outputNum(c chan<- int) {
	defer close(c)
	for i := 1; i <= 10; i++ {
		c <- i
	}
}

func getNum(c <-chan int) {
	defer wg3.Done()
	for i := range c {
		fmt.Println("读取到数值=", i)
	}
}

func main() {
	c := make(chan int)
	fmt.Println("begin...")
	wg3.Add(1)
	go outputNum(c)
	go getNum(c)
	wg3.Wait()
}
