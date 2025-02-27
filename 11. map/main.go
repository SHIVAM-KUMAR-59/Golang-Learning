package main

import "fmt"

// Map is a collection of key-value pairs where each key is unique. It can be unordered.
// A map is a reference type. When a map is assigned to a new variable, they both point to the same internal data structure.
// Therefore, changes made in one will reflect in the other.
// The key is used to retrieve the value.
// The key must be of a type that can be compared with the equality operator.
// The value can be of any type.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.

func main(){

	/*
	A map is created using the following syntax:
		var <mapName> map[<keyType>]<valueType>
	The keyType and valueType can be any type, including another map or slice.
	The keyType must be a comparable type. A type is comparable if it supports the == operator.
	The value can be of any type, including another map.
	*/

	var students map[string]int = map[string]int{
		"Alice": 90,
		"Bob": 80,
	}
	fmt.Println(students) // map[Alice:90 Bob:80]

	// Create a map using the make function
	// The make function returns a map of the given type, initialized and ready for use.
	// The make function takes two arguments: the map type and the capacity of the map.
	// We cannot initialize the map with values using the make function.
	// The capacity is the number of elements that the map can initially store.
	// The capacity is optional. If not provided, the map will be created with a default capacity.

	students_grade := make(map[string]int) // Create a map using the make function
	students_grade["Alice"] = 90
	students_grade["Bob"] = 80

	// Access the value using the key
	fmt.Println(students_grade["Alice"]) // 90
	fmt.Println(students_grade["Bob"]) // 80
	fmt.Println(students_grade) // map[Alice:90 Bob:80]

	// Update the value using the key
	students_grade["Alice"] = 95
	fmt.Println(students_grade) // map[Alice:95 Bob:80]

	// Delete a key-value pair from the map using the delete function
	delete(students_grade, "Bob")
	fmt.Println(students_grade) // map[Alice:95]

	// Check if a key exists in the map
	// The value, ok := map[key] syntax is used to check if a key exists in the map.
	grades, exists := students_grade["Alice"]
	if exists {
		fmt.Println("Alice's grade:", grades, "found", exists)
	} else {
		fmt.Println("Alice's grade not found")
	}

	// Iterate over the map
	// The range keyword is used to iterate over the map.
	for index, value := range students_grade {
		fmt.Println(index, value)
	}

}