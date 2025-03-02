package main

import (
	"fmt"
	"strconv"
)

// Data conversion: Converting one data type into another
// Rules of data conversion:
// 1. Conversion must be possible
// 2. Conversion must be explicit
// 3. Conversion must be safe
// 4. Conversion must be lossless

func main(){
	var num int = 42;
	fmt.Printf("Type of num is: %T\n", num) // %T gives the data type of variable

	// num += 1.23 This will throw an error because we are trying to add a float to an integer

	var data float64 = float64(num)
	fmt.Printf("Type of data is: %T\n", data)

	data += 1.23 // This will not throw an error because we are trying to add a float to a float

	// Converting to string
	var str string = strconv.Itoa(num) // strconv.Itoa() function is used to convert an integer to a string
	fmt.Printf("Type of str is: %T\n", str)

	// String to integer
	var num1 , _ = strconv.Atoi(str) // strconv.Atoi() function is used to convert a string to an integer
	fmt.Printf("Type of num1 is: %T\n", num1)

	// String to float
	var float_string string = "3.14"
	var float1 , _ = strconv.ParseFloat(float_string, 64) // strconv.ParseFloat() function is used to convert a string to a float
	fmt.Printf("Type of float1 is: %T\n", float1)
}