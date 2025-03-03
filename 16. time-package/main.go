package main

import (
	"fmt"
	"time"
)

// Time package in GoLang
// The time package provides functions for working with time values.
// The time package is used to represent and manipulate time values.
// Time Format: YYYY-MM-DD HR:MM:SS (24-hour format)

func main(){

	currentTime := time.Now() // Gives the current time
	fmt.Println(currentTime)
	fmt.Printf("Type of %T\n", currentTime) // time.Time time package provides its own data type

	formattedData := currentTime.Format("02-01-2006 15:04:05, Monday") // dd-mm-yyyy hh-mm-ss, day we have to write 02-01-2006
	fmt.Println("Formatted Time", formattedData)

	formattedData2 := currentTime.Format("2006/01/02, 3:04 PM") // 
	fmt.Println("Formatted Time 2", formattedData2)

	// Converting from string to time format
	layoutString := "2006-01-02" // yyyy-mm-dd the layout string decides the format of time returned
	dateString := "2025-03-03"
	timeFromDateString, _ := time.Parse(layoutString, dateString)
	fmt.Println("timeFromDateString", timeFromDateString)


}