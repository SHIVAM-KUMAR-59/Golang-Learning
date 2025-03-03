package main

import (
	"fmt"
	"time"
)

func main() {
	// Getting the current time
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Printf("Type of currentTime: %T\n", currentTime) // time.Time data type

	// Formatting time in different formats
	formattedData := currentTime.Format("02-01-2006 15:04:05, Monday")
	fmt.Println("Formatted Time:", formattedData)

	formattedData2 := currentTime.Format("2006/01/02, 3:04 PM")
	fmt.Println("Formatted Time 2:", formattedData2)

	// Converting a string to time format
	layoutString := "2006-01-02" // Layout string determines the format
	dateString := "2025-03-03"
	timeFromDateString, err := time.Parse(layoutString, dateString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed Time from String:", timeFromDateString)
	}

	// Adding and subtracting time
	oneWeekLater := currentTime.AddDate(0, 0, 7) // Adding 7 days
	fmt.Println("One week later:", oneWeekLater.Format("02-01-2006"))

	threeHoursBefore := currentTime.Add(-3 * time.Hour) // Subtracting 3 hours
	fmt.Println("Three hours ago:", threeHoursBefore.Format("02-01-2006 15:04:05"))

	// Calculating duration between two times
	pastTime := time.Date(2025, 3, 1, 12, 0, 0, 0, time.UTC)
	duration := currentTime.Sub(pastTime)
	fmt.Println("Duration since past time:", duration)

	// Extracting components from time
	fmt.Println("Year:", currentTime.Year())
	fmt.Println("Month:", currentTime.Month())
	fmt.Println("Day:", currentTime.Day())
	fmt.Println("Hour:", currentTime.Hour())
	fmt.Println("Minute:", currentTime.Minute())
	fmt.Println("Second:", currentTime.Second())

	// Checking if a date is before or after another date
	futureTime := currentTime.AddDate(0, 1, 0) // Adding 1 month
	fmt.Println("Is current time before future time?", currentTime.Before(futureTime))
	fmt.Println("Is current time after past time?", currentTime.After(pastTime))

	// Sleep function to pause execution for 2 seconds
	fmt.Println("Waiting for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("Resumed execution!")
}
