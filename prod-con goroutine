package main

import (
	"fmt"
	"sync"
)

func go1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i <= 10; i++ {
		ch <- i
	}
}

func go2(ch chan int, ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch1)
	for i := range ch {
		ch1 <- i + 1
	}
}

func go3(ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch1 {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go go1(ch, &wg)
	go go2(ch, ch1, &wg)
	go go3(ch1, &wg)
	wg.Wait()
}
