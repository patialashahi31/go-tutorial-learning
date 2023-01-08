package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	fmt.Println("Welcome to Go mod class")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":4000",r))
	
}

func greeter(){
	fmt.Println("Hey")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1> Hello </h1>"))
}







func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}
