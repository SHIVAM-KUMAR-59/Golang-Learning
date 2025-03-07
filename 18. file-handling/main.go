package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 1. Create a new file (or overwrite if it exists)
	file, err := os.Create("example.txt") // creates "example.txt"
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure file is closed when function exits

	// 2. Write content to the file
	content := "Hello World\nWelcome to Go File Handling!"
	_, writeErr := io.WriteString(file, content) // Writing string to file
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
		return
	}
	fmt.Println("File written successfully.")

	// 3. Read the file using Buffer
	openFile, openErr := os.Open("example.txt") // Open the file for reading
	if openErr != nil {
		fmt.Println("Error opening the file:", openErr)
		return
	}
	defer openFile.Close()

	fmt.Println("\nReading file using buffer:")
	buffer := make([]byte, 1024) // Create a buffer of 1024 bytes
	for {
		n, e := openFile.Read(buffer) // Read file into buffer
		if e == io.EOF {
			break // Stop reading at EOF
		}
		if e != nil {
			fmt.Println("Error reading from file:", e)
			return
		}
		fmt.Print(string(buffer[:n])) // Print the read content
	}

	// 4. Read the entire file using ioutil
	fmt.Println("\nReading file using ioutil:")
	data, readErr := ioutil.ReadFile("example.txt") // Read full file
	if readErr != nil {
		fmt.Println("Error reading file:", readErr)
		return
	}
	fmt.Println(string(data)) // Print file content

	// 5. Append new data to the file
	fmt.Println("\nAppending data to the file:")
	appendData := "\nThis is appended text!"
	appendFile, appendErr := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if appendErr != nil {
		fmt.Println("Error opening file for appending:", appendErr)
		return
	}
	defer appendFile.Close()

	_, appendWriteErr := appendFile.WriteString(appendData) // Append text
	if appendWriteErr != nil {
		fmt.Println("Error appending to file:", appendWriteErr)
		return
	}
	fmt.Println("Data appended successfully.")

	// 6. Read file using bufio scanner
	fmt.Println("\nReading file using bufio scanner:")
	readFile, readFileErr := os.Open("example.txt")
	if readFileErr != nil {
		fmt.Println("Error opening file:", readFileErr)
		return
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile) // Create a scanner
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Print line by line
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file line by line:", err)
	}
}
