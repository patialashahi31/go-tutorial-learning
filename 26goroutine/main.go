package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointer
var mut sync.Mutex // pointer

var signals = []string{"test"}

func main() {
	// go greeter("Hello")
	// greeter("World")
	websites := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _,web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}

	

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string){
// 	for i:=0; i<6; i++{
// 		time.Sleep(3* time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string){
	defer wg.Done()
	res , err := http.Get(endpoint)
	if err!= nil{
		fmt.Println("Not valid endpoint")
	}else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s \n",res.StatusCode,endpoint)
	}
}