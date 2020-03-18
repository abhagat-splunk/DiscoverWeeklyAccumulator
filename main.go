package main

import(
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	fmt.Println("Starting the application")
	response, err := http.Get("https://httpbin.org/ip")
	if err != nil{
		fmt.Printf("The HTTP request failed with error: %s\n", err)
	} else{
		data, _ := ioutil.ReadAll(response.body)
		fmt.Println(string(data))
	}

	jsonData := map[string]string{"firstname":"Ankit", "lastName":"Bhagat"}
	jsonValue := json.Marshal(jsonData)

	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil{
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else{
		data, _ = ioutil.ReadAll(response.body)
		fmt.Printf(string(data))
	}

	fmt.Println("Terminating the application.")
}