// mongodb+srv://patialashahi:<password>@cluster0.iwqknxd.mongodb.net/?retryWrites=true&w=majority
package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main(){
	fmt.Println("API with mongo db")
	r := router.Router()
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening to Port 4000")
}