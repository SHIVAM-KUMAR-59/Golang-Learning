package main

import (
	"fmt"     // Package for formatted I/O operations
	"net/url" // Package for URL parsing and modification
)

func main() {
	// Defining a raw URL string
	rawURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"

	// Parsing the URL string into a structured *url.URL object
	parsedURL, parseErr := url.Parse(rawURL)
	if parseErr != nil { // Checking if any error occurred during parsing
		fmt.Println("Error parsing URL:", parseErr)
		return // Exit the function if an error occurs
	}

	// Printing the type of parsedURL (should be *url.URL)
	fmt.Printf("Type of parsed URL: %T\n", parsedURL)

	// Printing the complete parsed URL
	fmt.Println("Parsed URL:", parsedURL)

	// Extracting and printing individual components of the URL
	fmt.Println("Scheme:", parsedURL.Scheme)   // Output: https (protocol used)
	fmt.Println("Host:", parsedURL.Host)       // Output: example.com:8080 (domain with port)
	fmt.Println("Path:", parsedURL.Path)       // Output: /path/to/resource (path of resource)
	fmt.Println("Raw Query:", parsedURL.RawQuery) // Output: key1=value1&key2=value2 (query parameters)

	// Modifying the URL components
	parsedURL.Path = "/newPath"        // Changing the path from "/path/to/resource" to "/newPath"
	parsedURL.RawQuery = "username=abc" // Changing the query parameters

	// Constructing a new URL string from the modified URL object
	newURL := parsedURL.String()

	// Printing the new URL after modifications
	fmt.Println("New URL:", newURL) // Output: https://example.com:8080/newPath?username=abc
}
