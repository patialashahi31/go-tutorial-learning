package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in go ")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	// Receive only
	go func(ch <-chan int, wg *sync.WaitGroup){
		val , isChannelOpen := <-myCh
		fmt.Println(val,isChannelOpen)
		// fmt.Println(<-myCh)
		wg.Done()

	}(myCh,wg)

	// Send only
	go func(ch chan<- int, wg *sync.WaitGroup){
		
		// myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh,wg)

	wg.Wait()

	
}