package main // Every Go program must start with a package declaration. The main package is used to create an executable file.
// An executable program must contain a main package.

import "fmt" // fmt is a standard library package that provides formatted I/O functions.
// Go uses relative imports to bring packages into the file.

func main(){
	fmt.Println("Hello World")
}