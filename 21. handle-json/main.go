package main

import (
	"encoding/json" // Package for working with JSON data
	"fmt"           // Package for formatted I/O operations
)

// Person struct represents a user with name, age, and adult status
type Person struct {
	Name    string `json:"name"`    // Struct field mapped to "name" in JSON
	Age     int    `json:"age"`     // Struct field mapped to "age" in JSON
	IsAdult bool   `json:"isAdult"` // Struct field mapped to "isAdult" in JSON
}

func main() {
	// Creating an instance of Person struct
	person := Person{Name: "ABC", Age: 20, IsAdult: true}
	fmt.Println("Person data:", person)

	// Converting (Encoding/Marshalling) struct into JSON format
	jsonData, marshalErr := json.Marshal(person)
	if marshalErr != nil { // Handling encoding errors
		fmt.Println("Error marshalling:", marshalErr)
		return
	}

	// Printing the JSON representation of the struct
	fmt.Println("JSON Data:", string(jsonData)) // Output: {"name":"ABC","age":20,"isAdult":true}

	// Decoding (Unmarshalling) JSON back into struct
	var personData Person
	unmarshalErr := json.Unmarshal(jsonData, &personData)
	if unmarshalErr != nil { // Handling decoding errors
		fmt.Println("Error Unmarshalling:", unmarshalErr)
		return
	}

	// Printing the struct obtained after unmarshalling
	fmt.Println("Unmarshalled data:", personData) // Output: {ABC 20 true}
}
