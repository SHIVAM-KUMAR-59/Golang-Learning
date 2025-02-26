package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("What is your name?");
	var name string;
	fmt.Scan(&name); // Scan reads the input from the user till the first whitespace and stores it in the variable
	fmt.Println("Hello,", name);

	fmt.Println("What is your full name?");
	reader := bufio.NewReader(os.Stdin); // Create a reader object of buffer reader to read input from the user
	fullName, _ := reader.ReadString('\n'); // Read till '\n' is encountered
	fmt.Println(fullName)
}