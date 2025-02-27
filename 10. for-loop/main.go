package main

import "fmt"

// There is no while loop or do-while loop in Go, but we can use for loop to achieve the same

func main(){
	// for loop syntax
	/*
		for initialization; condition; post {
			// code
		}
	*/

	for i := 0; i < 5; i++ {
		fmt.Println("Number is:", i)
	}

	// Infinite loop using for loop and break statement
	count := 0
	for {
		fmt.Println("Inside infinite loop")
		count++
		if count == 5 {
			break
		}
	}

	// Range in for loop
	/*
		range keyword is used to iterate over items in a variety of data structures like arrays, slices, maps, strings, etc.
		Sytanx:
		for index, value := range data {
			// code
		}
	*/
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Println("Index:", index, "Value:", value)
	}
}