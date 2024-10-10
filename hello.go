package main

import "fmt"

func main() {
    fmt.Println("Hello, World! This is my first code of Golang.")
    
    message := Hello()
    fmt.Println(message)

    // Declare and initialize the variable 'a'
    a := 23
    fmt.Println(a)
}

// Hello function that returns a string
func Hello() string {
    return "Hello, World! This is my first code of Golang."
}
