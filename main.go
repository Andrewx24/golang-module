// main.go
package main

import (
  "fmt"
)

// Defining a struct for a person
type Person struct {
    Name           string
    Age            int
    Occupation     string
    MarriageStatus bool
}

func main() {
   // Initializing a variable of type Person
   person := Person{
       Name:           "John Doe",
       Age:            25,
       Occupation:     "Gardener",
       MarriageStatus: false,
   }

   // Printing the values
   fmt.Println("Hello, World!")
   fmt.Println("a:", 10)
   
   // Printing the struct's fields
   fmt.Println("Person details:")
   fmt.Printf("Name: %s\n", person.Name)
   fmt.Printf("Age: %d\n", person.Age)
   fmt.Printf("Occupation: %s\n", person.Occupation)
   fmt.Printf("Marriage Status: %t\n", person.MarriageStatus)
}
