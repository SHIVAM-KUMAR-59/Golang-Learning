package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Todo struct represents the structure of a TODO item
// fetched from the JSONPlaceholder API.
type Todo struct {
	UserId    int    `json:"userId"`    // User ID of the TODO item
	Id        int    `json:"id"`        // Unique ID of the TODO item
	Title     string `json:"title"`     // Title/description of the TODO item
	Completed bool   `json:"completed"` // Status of completion (true/false)
}

func performGetMethod(){
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

func performPostMethod(){
	todo := Todo{ // Create a new todo to send as data
		UserId: 23,
		Title: "My First Todo",
		Completed: true,
	}

	// Convert the Todo struct to JSON
	jsonData, marshalErr := json.Marshal(todo)
	if marshalErr != nil {
		fmt.Println("Error marshalling:", marshalErr)
		return
	}

	// Convert JSON data into string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	// Send a POST request to create a new TODO item
	res, postErr := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if postErr != nil{
		fmt.Println("Error posting data", postErr)
		return
	}
	defer res.Body.Close()

	// Check if the response status code is not 201 (Created)
	if res.StatusCode != http.StatusCreated {
		fmt.Println("Error getting data:", res.StatusCode)
		return
	}

	// Convert the response in readable from
	data, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil{
		fmt.Println("Error while reading response:", readErr)
		return
	}

	fmt.Println(string(data))

}

func performUpdateMethod(){
	todo := Todo{ // Create a new todo to send as data
		UserId: 239,
		Title: "Learning GoLang",
		Completed: false,
	}

	// Convert the Todo struct to JSON
	jsonData, marshalErr := json.Marshal(todo)
	if marshalErr != nil {
		fmt.Println("Error marshalling:", marshalErr)
		return
	}

	// Convert JSON data into string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	// Create a PUT request to create a new TODO item
	req, putErr := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", jsonReader)
	if putErr != nil{
		fmt.Println("Error Updating data", putErr)
		return
	}
	// Set the content type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	res, sendErr := http.DefaultClient.Do(req)
	if sendErr != nil{
		fmt.Println("Error sending data", sendErr)
		return
	}
	defer res.Body.Close()

	// Check if the response status code is not 200 (OK)
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error getting data:", res.StatusCode)
		return
	}

	// Convert the response in readable from
	data, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil{
		fmt.Println("Error while reading response:", readErr)
		return
	}

	fmt.Println(string(data))
}

func performDeleteMethod(){

	// Create DELETE Request
	req, delErr := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if delErr != nil{
		fmt.Println("Error Deleting Data", delErr)
	}

	// Send the request
	res, sendErr := http.DefaultClient.Do(req)
	if sendErr != nil{
		fmt.Println("Error sending data", sendErr)
		return
	}
	defer res.Body.Close()

	// Check the status code
	if(res.StatusCode != http.StatusOK){
		fmt.Println("Error getting data:", res.StatusCode)
		return
	}

	// Read the response body
	data, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil{
		fmt.Println("Error while reading response:", readErr)
		return
	}

	fmt.Println(string(data))
	fmt.Println("Deleted Successfully with status code: ", res.StatusCode )
}

func main() {
	
	// performGetMethod()
	// performPostMethod()
	// performUpdateMethod()
	// performDeleteMethod()
}
