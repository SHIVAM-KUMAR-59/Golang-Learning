package main

import (
	"fmt"
	"sync"
)

/*
sync.WaitGroup
--------------------------
- `sync.WaitGroup` is used to synchronize multiple goroutines.
- It allows the main function to wait until all goroutines complete execution.
- `WaitGroup` prevents the main function from exiting before goroutines finish.
- `Add()`, `Done()`, and `Wait()` are the three main methods:

1. `Add(n)`: Adds `n` to the counter to track active goroutines.
2. `Done()`: Decrements the counter when a goroutine completes.
3. `Wait()`: Blocks execution until the counter becomes zero.

🔹 Example Analogy:
- Imagine a teacher (main function) assigning tasks to students (goroutines).
- The teacher keeps track of students working (`Add()`).
- Each student reports completion (`Done()`).
- The teacher waits (`Wait()`) until all students finish before ending class.
*/

// worker function simulates a worker performing a task
func worker(i int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done() // Decrements the counter when this function completes

	fmt.Printf("Worker %d Started Working\n", i+1)

	// Simulate work being done (can be replaced with actual computation)
	fmt.Printf("Worker %d End Working\n", i+1)
}

func main() {
	var waitGroup sync.WaitGroup // Declare a WaitGroup to synchronize goroutines

	// Start 3 worker goroutines
	for i := 0; i < 3; i++ {
		waitGroup.Add(1) // Increments the counter before starting a new goroutine
		go worker(i, &waitGroup) // Launch goroutine to perform work
	}

	waitGroup.Wait() // Blocks until all worker goroutines call Done()
	fmt.Println("Workers Task Completed") // Executes only after all workers finish
}
