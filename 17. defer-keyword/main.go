package main

import "fmt"

// Defer Keyword
// - Defer statements are executed in Last-In-First-Out (LIFO) order (like a stack).
// - Used for cleanup tasks, closing files, unlocking resources, etc.

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println("Starting of Program")

	data := add(2, 3)
	defer fmt.Println("Deferred: Data", data) // This gets executed at the end

	defer fmt.Println("Deferred: Middle of Program") // This will execute second-last

	fmt.Println("End of Program")

	// Example 1: Multiple defer statements (LIFO execution)
	defer fmt.Println("Deferred: First")
	defer fmt.Println("Deferred: Second")
	defer fmt.Println("Deferred: Third") // This will execute first among the deferred calls

	// Example 2: Defer with function call
	defer func() {
		fmt.Println("Deferred: Inside an anonymous function")
	}()

	// Example 3: Defer with arguments evaluation
	num1, num2 := 10, 5
	defer fmt.Println("Deferred: Addition result:", add(num1, num2)) // Arguments are evaluated immediately
	num1 = 20 // Changing value, but it won't affect the deferred function

	// Example 4: Defer inside a function
	deferExample()

	// Example 5: Defer with resource cleanup
	openResource()
}

// Function demonstrating defer execution order
func deferExample() {
	fmt.Println("Inside deferExample function")
	defer fmt.Println("Deferred: deferExample - First")
	defer fmt.Println("Deferred: deferExample - Second")
	fmt.Println("Exiting deferExample function")
}

// Simulating opening and closing a resource
func openResource() {
	fmt.Println("Opening a resource (e.g., file, database)")
	defer fmt.Println("Closing the resource (deferred cleanup)") // Ensuring cleanup happens
	fmt.Println("Using the resource")
}
