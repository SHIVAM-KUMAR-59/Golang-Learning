package main

import "fmt"

// Slices are a better version of arrays, it is a flexible version of arrays
// Slices are like references to arrays
// A slice does not store any data, it just describes a section of an underlying array
// Changing the elements of a slice modifies the corresponding elements of its underlying array
// Other slices that share the same underlying array will see those changes

func main(){

	// Declaring a slice
	// A slice is declared by specifying the type of the elements it will hold followed by square brackets []
	// The number of elements in the slice is not specified
	// A slice is created using the built-in make function

	// The length of the slice is the number of elements the slice contains
	// The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created
	// The length and capacity of a slice can be obtained using the built-in len and cap functions respectively

	numbers := []int{1, 2, 3, 4, 5};
	numbers = append(numbers, 6, 7, 8, 9, 10); // Append elements to the slice
	fmt.Println(numbers);
	fmt.Println(len(numbers)); // Length
	fmt.Println(cap(numbers)); // Capacity

	// Use the make function to create a slice
	// The make function creates an array and returns a slice that refers to that array
	// The make function takes three arguments: the type of the slice, the length of the slice, and the capacity of the slice
	// make([]T, length, capacity)
	// The capacity argument is optional and defaults to the length of the slice
    arr := make([]int, 5, 5);
	fmt.Println(arr);
	arr = append(arr, 6, 7, 8, 9, 10); // Append elements to the slice arr starting from 6th after the last capacity is reached the capacity is doubled 
	fmt.Println(arr);

	// Way to define an empty slice
	var emptySlice []int;
	fmt.Println(emptySlice);

	// Define an empty slice using the make function
	emptySlice = make([]int, 0); // The length and capacity are 0
	fmt.Println(emptySlice);

	// Accessing elements of a slice
	fmt.Println("Element at index 0 is", arr[0]);
}