package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string `json:"-"` // remove from json
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json class")
	
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	lcoCourses := []course{
		{"React", 299, "Learn.online","abc",[]string{"web","js"}},
		{"Python", 199, "Learn.online","13",nil},
		{"Ruby", 399, "Learn.online","a5",[]string{"web","fullstack"}},
	}

	finalJson,err := json.MarshalIndent(lcoCourses,"","\t")
	checkNilError(err)
	fmt.Printf("%s\n",finalJson)

}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}
}
	







func checkNilError(err error){
	if err != nil {
		panic(err)
	}
}
