package main

import "fmt"

func main(){

	// Switch statement syntax
	/*
		switch <initialization>; <condition> {
			case <value1>:
				<code>
			case <value2>:
				<code>
			default: (If no case matches)
				<code>
		}
	*/

	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:	
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Multiple expressions in case
	/*
		switch <initialization>; <condition> {
			case <value1>, <value2>, <value3>:
				<code>
			default:
				<code>
		}
	*/

	month := 2
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 days")
	case 4, 6, 9, 11:
		fmt.Println("30 days")
	case 2:
		fmt.Println("28/29 days")
	default:
		fmt.Println("Invalid month")
	}
}