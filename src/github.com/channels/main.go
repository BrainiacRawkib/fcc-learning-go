package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) {
		// receive only channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		// send only channel
		ch <- 42
		wg.Done()
	}(ch)

	// Buffered Channels
	ch2 := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		// receive only channel
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}(ch2)
	go func(ch chan<- int) {
		// send only channel
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch2)

	// for-range in channels
	ch3 := make(chan int, 50)
	wg.Add(2)
	// receive only channel
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch3)
	// send only channel
	go func(ch chan<- int) {
		ch <- 42
		close(ch)
		ch <- 27
		close(ch)
		wg.Done()
	}(ch3)

	//go func() {
	//	i := <-ch
	//	fmt.Println(i)
	//	wg.Done()
	//}()
	//
	//for j := 0; j < 5; j++ {
	//	wg.Add(2)
	//	go func() {
	//		ch <- 42
	//		wg.Done()
	//	}()
	//}
	wg.Wait()
}
