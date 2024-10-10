package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))
    secretNumber := randomSource.Intn(100) + 1
    var guess int

    fmt.Println("Welcome to the Number Guessing Game!")
    fmt.Println("I'm thinking of a number between 1 and 100. Can you guess it?")

    for attempts := 0; attempts < 10; attempts++ {
        fmt.Print("Enter your guess: ")
        fmt.Scan(&guess)

        if guess < secretNumber {
            fmt.Println("Too low! Try again.")
        } else if guess > secretNumber {
            fmt.Println("Too high! Try again.")
        } else {
            fmt.Printf("Congratulations! You guessed the number %d!\n", secretNumber)
            return
        }
    }
    fmt.Printf("Out of attempts! The number was %d.\n", secretNumber)
}
