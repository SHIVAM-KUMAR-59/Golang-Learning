package main

import "fmt"

// An array is a collection of elements of the same data type.
// The elements of an array are stored in contiguous memory locations.
// The size of an array is fixed and cannot be changed once it is declared.
// All the elements of the array/slice are initialized to their zero value by default.
// The zero value of an integer is 0, of a float is 0.0, of a string is "", and of a boolean is false.

func main(){
	// var array_name[size] data_type
	var arr[5] int;
	
	fmt.Println(arr);

	// Another way to declare and initialize an array
	// var array_name[size] data_type = [size]data_type{values}
	var arr2[5] int = [5]int{10, 20, 30, 40, 50};
	fmt.Println(arr2);

	var arr3 = [5]int{10, 20, 30, 40, 50}; // The size of the array is inferred from the number of values
	fmt.Println(arr3);

	arr4 := [5]int{10, 20, 30, 40, 50}; // The var keyword can be omitted using the short declaration syntax
	fmt.Println(arr4);

	// Accessing elements of an array
	fmt.Println("Element at index 0 is", arr4[0]);
	fmt.Println("Element at index 1 is", arr4[1]);
	fmt.Println("Element at index 2 is", arr4[2]);
	fmt.Println("Element at index 3 is", arr4[3]);
	fmt.Println("Element at index 4 is", arr4[4]);
	// fmt.Println("Element at index 4 is", arr4[4]); This will give error because the index is out of range

	// Getting the length of the array
	fmt.Println("Length og arr4 is", len(arr4));

}