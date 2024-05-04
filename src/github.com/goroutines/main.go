package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// sayMessage
	fmt.Printf("\n")
	go sayMessage("Hello Go!")

	var msg string = "Hello from first!"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodie Goodie!" // this gets printed out because of the race condition
	//time.Sleep(100 * time.Millisecond)

	// using closures to decouple the main function from the goroutine
	var msg1 string = "Hello from second!"
	go func(msg string) {
		fmt.Println(msg1)
	}(msg1)
	msg1 = "Goodie Goodie 2!"
	//time.Sleep(100 * time.Millisecond)

	// using WaitGroup to decouple the main function from the goroutine
	var msg2 string = "Hello from third!"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg2)
		wg.Done()
	}(msg2)
	msg2 = "Goodie Goodie 3!"
	wg.Wait()
	time.Sleep(100 * time.Millisecond)

	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayIncrement()
		m.Lock()
		go increment()
	}
	wg.Wait()

	fmt.Printf("Threads: %v", runtime.GOMAXPROCS(100))
}

func sayMessage(msg string) {
	fmt.Printf("*** sayMessage ***\n")
	fmt.Println(msg)
	i := 3
	fmt.Println(i)
}

func sayIncrement() {
	fmt.Printf("Incrementing #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
