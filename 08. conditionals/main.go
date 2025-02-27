package main

import "fmt"

func main(){
	/*
	Conditionals: if-else statement
	- if condition {
		// code
	} else {
		// code
	}	
	*/

	x := 10
	if x == 10 {
		fmt.Println("x is 10")
	} else {
		fmt.Println("x is not 10")
	}

	/*
	Nested if-else ladder
	- if condition {
		// code
	} else if condition {
	    // code
	} else {
	 	// code
	}
	*/

	if x == 10 {
		fmt.Println("x is 10")
	} else if x > 10 {
		fmt.Println("x is greater than 10")
	}else{
		fmt.Println("x is less than 10")
	}

	if x > 10 && (x < 20 || x == 30) {
		fmt.Println("x is between 10 and 20 or x is 30")
	}
}