package main

/*
GoRoutines
--------------------------
Ways to execute a program:
1. Sequential
2. Concurrent or Parallelism

Concurrncy: Multiple tasks running at the same time
Parallelism: Multiple processes running at the same time

Concurrency vs. Parallelism:
- Concurrency: Multiple tasks making progress at the same time (task switching).
- Parallelism: Multiple tasks executing at the exact same time using multiple CPU cores.

Example Analogy:
- Concurrency: A single chef cooking multiple dishes by switching between them.
- Parallelism: Multiple chefs cooking different dishes at the same time.

Goroutines:
- A goroutine is a lightweight thread of execution in Go.
- It allows running multiple functions concurrently without creating heavy OS threads.
- Goroutines are managed by the Go runtime, making them efficient and scalable.

How to create a goroutine?
Use the `go` keyword before a function call to execute it concurrently.

*/

import "fmt"

func main(){
	fmt.Println("Learning Go-Routines")
}