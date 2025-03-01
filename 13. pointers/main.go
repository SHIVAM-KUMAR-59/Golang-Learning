package main

import "fmt"

// Pointers in GoLang
// A pointer holds the memory address of a value. The type *T is a pointer to a T value. Its zero value is nil.
// Properties of Pointers in GoLang:
// 1. A pointer is a variable that stores the memory address of another variable.
// 2. A pointer is declared using the * operator.
// 3. A pointer is used to share the memory address of a variable.

func main(){

	// Declaring a pointer
	var ptr *int // ptr is a pointer to an integer
	ptr = new(int) // new() function is used to allocate memory to a pointer variable

	// Assigning a value to a pointer
	*ptr = 100 // 100 is assigned to the memory address stored in ptr

	fmt.Println("Value of ptr:", *ptr) // Output: 100

	var ptr1 *int
	var num int = 5;
	ptr1 = &num // & operator is used to get the memory address of a variable
	fmt.Println("Value of ptr1:", *ptr1) // Output: 5
}