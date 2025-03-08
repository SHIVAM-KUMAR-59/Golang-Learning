package main

import (
	"encoding/json" // Package for encoding and decoding JSON
	"fmt"           // Package for formatted I/O
	"net/http"      // Package for HTTP client and server implementation
)

// Todo struct represents the structure of a TODO item
// fetched from the JSONPlaceholder API.
type Todo struct {
	UserId    int    `json:"userId"`    // User ID of the TODO item
	Id        int    `json:"id"`        // Unique ID of the TODO item
	Title     string `json:"title"`     // Title/description of the TODO item
	Completed bool   `json:"completed"` // Status of completion (true/false)
}

func main() {
	// Send a GET request to fetch TODO item with ID 1
	res, getErr := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if getErr != nil {
		// Print error message if the request fails
		fmt.Println("Error getting data:", getErr)
		return
	}

	// Ensure that the response body is closed after function execution
	defer res.Body.Close()

	// Check if the response status code is not 200 (OK)
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error getting data:", res.StatusCode)
		return
	}

	// Uncomment this section if you want to manually read and print the response body
	/*
		data, ioErr := ioutil.ReadAll(res.Body)
		if ioErr != nil {
			fmt.Println("Error while reading response:", ioErr)
			return
		}
		fmt.Println("Response is:", string(data))
	*/

	// Create a variable to store the decoded JSON response
	var todo Todo

	// Decode the JSON response into the Todo struct
	decErr := json.NewDecoder(res.Body).Decode(&todo)
	if decErr != nil {
		fmt.Println("Error decoding data:", decErr)
		return
	}

	// Print the parsed TODO item
	fmt.Println("Response is:", todo)
}
