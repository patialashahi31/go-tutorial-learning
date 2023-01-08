package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race conditions")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	
	var score = []int{0}
	

	wg.Add(4) // 4 go routines
	go func(wg *sync.WaitGroup, mut *sync.RWMutex){
		fmt.Println("First")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex){
		fmt.Println("Second")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex){
		fmt.Println("third")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex){
		fmt.Println("third")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg,mut)

	wg.Wait()

	fmt.Println(score)
}