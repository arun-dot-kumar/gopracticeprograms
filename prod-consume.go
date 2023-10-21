package main

import (
	"fmt"
	"sync"
	"time"
)

func produce(ch chan int, ch1 chan struct{}, wg *sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		ch1 <- struct{}{}
		ch <- i
		fmt.Printf("%d is sending", i)
		fmt.Println()
	}
	close(ch)
	close(ch1)
	wg.Done()
}

func consume(ch chan int, ch1 chan struct{}, wg *sync.WaitGroup) {
	for i := range ch {
		go func(ch1 chan struct{}, i int) {
			time.Sleep(2 * time.Second)
			fmt.Println(i)
			<-ch1
		}(ch1, i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan struct{}, 10)
	ch := make(chan int)
	go produce(ch, ch1, &wg)
	go consume(ch, ch1, &wg)
	wg.Add(2)
	wg.Wait()
}
