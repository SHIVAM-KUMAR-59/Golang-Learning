package main

import "fmt"

// divide function takes two float64 numbers as input and returns
// the result of their division along with an error (if any).
func divide(a, b float64) (float64, error) {
    // Check if the denominator is zero to prevent division by zero error.
    if b == 0 {
        return 0, fmt.Errorf("Undefined") // Return an error message if b is zero.
    } else {
        return a / b, nil // Return the division result and nil error if valid.
    }
}

func main() {
    fmt.Println("---------------------- Error Handling ---------------------")
    
    // Call the divide function with 10 and 0 as arguments.
    // The second return value (error) is ignored using `_`. `_` is a blank identifier in Go.
    ans, _ := divide(10, 0)
    
    // Print the result of division.
    // Since we ignored the error, it will print 0 (default return value in case of error).
    fmt.Println(ans)
}
