package main

import (
	"fmt"
	"strings" // strings package
)

// Strings in go are managed by the Go runtime. It is a collection of bytes. The zero value of a string is an empty string.
// There is a String package in the Go standard library that provides functions for working with strings.

func main(){

	data := "apple,orange,banana"
	parts := strings.Split(data, ",") // Split() function is used to split a string into a slice of substrings
	fmt.Println(parts)

	count := strings.Count(data, "a") // Count() function is used to count the number of occurrences of a substring in a string
	fmt.Println(count)

	str := "       Hello World       "
	trimmed := strings.TrimSpace(str) // TrimSpace() function is used to remove leading and trailing whitespace
	fmt.Println(trimmed)

	result := strings.Join([]string{data, trimmed}, " ") // Join() function is used to join a slice of strings into a single string
	fmt.Println(result)

}