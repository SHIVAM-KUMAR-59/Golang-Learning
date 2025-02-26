package main

import "fmt"

func main(){
	age :=  25;
	name := "John Doe";
	weight := 75.5;

	fmt.Println("Name: ", name, "Age: ", age, "Weight: ", weight); // Adds space between each value and inserts a new line at the end by default
	fmt.Print("Name: ", name, "Age: ", age, "Weight: ", weight, "\n"); // Does not add space between each value and does not insert a new line at the end
	fmt.Printf("Name: %s Age: %d Weight: %f\n", name, age, weight); // Uses format specifiers to format the output
}