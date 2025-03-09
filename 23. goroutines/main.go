package main

/*
GoRoutines
--------------------------
Ways to execute a program:
1. Sequential - Tasks execute one after another.
2. Concurrent - Multiple tasks progress at the same time.
3. Parallel - Multiple tasks execute at the exact same time using multiple CPU cores.

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
- Use the `go` keyword before a function call to execute it concurrently.

Important Notes:
- The `main` function does NOT wait for goroutines to finish execution by default.
- If the main function exits, all running goroutines are also terminated.
- To prevent premature termination, we can use `time.Sleep()` or proper synchronization methods like `sync.WaitGroup`.

*/

import (
	"fmt"
	"time"
)

// sayHello prints a message, waits for 2 seconds, and then prints another message
func sayHello() {
	fmt.Println("sayHello() Function Started")
	time.Sleep(2000 * time.Millisecond) // Simulating work with a delay
	fmt.Println("sayHello() Function Ended")
}

// sayHi prints a message, waits for 1 second, and then prints another message
func sayHi() {
	fmt.Println("sayHi() Function Started")
	time.Sleep(1000 * time.Millisecond) // Simulating work with a delay
	fmt.Println("sayHi() Function Ended")
}

func main() {
	// Calling both functions concurrently using goroutines
	go sayHello()
	go sayHi()

	// 🟡 Without a delay, the main function will exit before goroutines finish
	// Waiting to allow goroutines to complete their execution
	time.Sleep(4000 * time.Millisecond) // Giving enough time for goroutines to finish

/*
Expected Output (Non-Deterministic Order Due to Concurrency):
--------------------------------------------------------------
sayHello() Function Started
sayHi() Function Started
sayHi() Function Ended
sayHello() Function Ended

Why does `sayHi()` finish first?
- `sayHi()` sleeps for 1 second.
- `sayHello()` sleeps for 2 seconds.
- Since `sayHi()` has a shorter delay, it completes first.

*/
}
