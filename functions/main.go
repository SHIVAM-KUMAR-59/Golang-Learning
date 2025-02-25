package main

import "fmt"

// Function
/*
func function_name(parameter_a, parameter_b, ..., parameter_n) (return_type_1, return_type_2, ...., return_type_n){
	function body
}
*/

func sum(a int, b int) (int, string){
	fmt.Println("Inside function")
	return a + b, "Sum"
}

// This is another way to define a function, here a variable result is returned with the calculated value
func add(a int, b int) (result int){
	return a + b
}

func main(){
	var c, d  = sum(10, 20);
	fmt.Println(c)
	fmt.Println(d)

	var a = add(10, 20);
	fmt.Println(a)
}