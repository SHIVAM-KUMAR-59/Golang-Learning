package main

import (
	"fmt"       // Package for formatted I/O operations
	"io/ioutil" // Package for reading from response body
	"net/http"  // Package for making HTTP requests
)

func main() {
	fmt.Println("Learning Web Requests in GoLang")

	// Making a GET request to fetch data from the given URL
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil { // Checking if there was an error in making the request
		fmt.Println("Error getting data:", err)
		return // Exit the function if an error occurs
	}
	// Ensure that the response body is closed once the function exits
	defer res.Body.Close()

	// Reading the response body into a byte slice
	data, ioErr := ioutil.ReadAll(res.Body)
	if ioErr != nil { // Checking for errors while reading response data
		fmt.Println("Error while reading response:", ioErr)
		return // Exit the function if an error occurs
	}

	// Converting the byte slice to a string and printing the response
	fmt.Println("Response is:", string(data))
}
