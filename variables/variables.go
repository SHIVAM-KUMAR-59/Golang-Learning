package main

import "fmt"

func main(){
	// var is a keyword used to declare a variable. name is the name of the variable, and string is the type of the variable.
	var name string = "Alex"; // Adding the type of the variable is optional. Go can infer the type of the variable, but if specified no other type can be assigned to it

	// var name string  = 89 // This will throw an error because the variable name is of type string and we are trying to assign an integer to it

	var version = "1.0" // Here we are not specifying the type of the variable, Go will infer the type of the variable

	var money int = 100;

	// Go also allows you to declare multiple variables in a single line
	var (
		age int = 20
		cgpa float32 = 3.5
		passed bool = true
	)

	fmt.Println(age)
	fmt.Println(cgpa)
	fmt.Println(passed)
	fmt.Println(money)
	fmt.Println(name)
	fmt.Println(version)

	const pi float64 = 3.14 // const is a keyword used to declare a constant. Constants are variables whose values cannot be changed once they are assigned.
	// pi = 3.14159 // This will throw an error because we are trying to change the value of a constant
	fmt.Println(pi)
}