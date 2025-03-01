package main

import "fmt"

// Structures (structs) in Go are user-defined data types that allow us to combine
// data items of different types. Structs are useful for representing records.

// Defining a struct named 'Person'
type Person struct {
	firstName string // Field or member variable
	lastName  string
	age       int
}

// Defining a struct named 'Contact' to store email and phone details
type Contact struct {
	email string
	phone string
}

// Defining a struct named 'Address' to store address-related details
type Address struct {
	houseNumber int
	area        string
	city        string
}

// Defining a struct named 'Employee' which embeds other structs
type Employee struct {
	// Nested struct (composition) to store multiple details about an employee
	Person_Details  Person  // Embedding 'Person' struct
	Contact_Details Contact // Embedding 'Contact' struct
	Address_Details Address // Embedding 'Address' struct
}

func main() {
	// Creating a struct instance (emp1) and initializing it with values
	emp1 := Person{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
	}

	// Accessing struct fields using dot notation
	fmt.Println("First Name:", emp1.firstName)
	fmt.Println("Last Name:", emp1.lastName)
	fmt.Println("Age:", emp1.age)

	// Using 'new' keyword to create a struct pointer
	emp2 := new(Person) // Returns a pointer to the struct (heap allocation)
	emp2.firstName = "Jack"
	emp2.lastName = "Adams"
	emp2.age = 25

	// Printing the struct pointer and its fields
	fmt.Println("Emp2:", emp2) // &{Jack Adams 25}, '&' indicates a pointer to the struct
	fmt.Println("First Name:", emp2.firstName)
	fmt.Println("Last Name:", emp2.lastName)
	fmt.Println("Age:", emp2.age)

	// Creating an instance of Employee struct with nested struct fields
	employee := Employee{ // 'employee' variable follows camelCase naming convention
		Person_Details: Person{
			firstName: "John",
			lastName:  "Doe",
			age:       25,
		},
		Contact_Details: Contact{
			email: "abc@gmail.com",
			phone: "1234567890",
		},
		Address_Details: Address{
			houseNumber: 123,
			area:        "XYZ",
			city:        "ABC",
		},
	}

	// Printing the nested struct details
	fmt.Println("Employee Details:", employee)
	fmt.Println("Personal Details:", employee.Person_Details)
	fmt.Println("Contact Details:", employee.Contact_Details)
	fmt.Println("Address Details:", employee.Address_Details)
}
