package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	fmt.Println("Learning Web Requests in GoLang")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil{
		fmt.Println("Error getting data", err)
		return
	}
	defer res.Body.Close()

	data, ioErr := ioutil.ReadAll(res.Body)
	if(ioErr) != nil{
		fmt.Println("Error while reading response", ioErr)
		return
	}

	fmt.Println("Resoponse is:", string(data))
}